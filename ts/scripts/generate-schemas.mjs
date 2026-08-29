#!/usr/bin/env node
/**
 * Generate Zod schemas from the SAS JSON Schema.
 *
 * The Go structs in package sas are the source of truth. This script
 * converts the generated JSON Schema (schema/architecture.schema.json)
 * into Zod schemas + TypeScript types so the TypeScript implementation
 * stays downstream of the Go model rather than a second source of truth.
 *
 * Usage: node scripts/generate-schemas.mjs
 */

import { readFileSync, writeFileSync, mkdirSync } from 'fs';
import { dirname, join, resolve } from 'path';
import { fileURLToPath } from 'url';
import { jsonSchemaToZod } from 'json-schema-to-zod';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const SCHEMA_DIR = resolve(__dirname, '../../schema');
const OUTPUT_DIR = resolve(__dirname, '../src/generated');

const SCHEMAS = [{ file: 'architecture.schema.json', name: 'Architecture' }];

/**
 * Resolve $ref references in a JSON Schema by inlining $defs, so each
 * generated Zod definition is self-contained.
 */
function resolveRefs(schema, defs) {
  if (typeof schema !== 'object' || schema === null) {
    return schema;
  }
  if (typeof schema === 'boolean') {
    return schema ? {} : { not: {} };
  }
  if (schema.$ref) {
    const refPath = schema.$ref;
    if (refPath.startsWith('#/$defs/')) {
      const defName = refPath.slice(8);
      if (defs[defName]) {
        return resolveRefs(defs[defName], defs);
      }
    }
    return schema;
  }

  const resolved = {};
  for (const [key, value] of Object.entries(schema)) {
    if (key === '$defs' || key === 'definitions') {
      continue;
    }
    if (Array.isArray(value)) {
      resolved[key] = value.map((item) => resolveRefs(item, defs));
    } else if (typeof value === 'object' && value !== null) {
      resolved[key] = resolveRefs(value, defs);
    } else {
      resolved[key] = value;
    }
  }
  return resolved;
}

/**
 * Topological sort of definitions based on their $ref dependencies, for
 * readable output ordering.
 */
function topologicalSort(defs) {
  const visited = new Set();
  const result = [];

  function getDeps(schema) {
    const deps = [];
    const traverse = (obj) => {
      if (typeof obj !== 'object' || obj === null) return;
      if (obj.$ref && obj.$ref.startsWith('#/$defs/')) {
        deps.push(obj.$ref.slice(8));
      }
      for (const value of Object.values(obj)) {
        traverse(value);
      }
    };
    traverse(schema);
    return deps;
  }

  function visit(name) {
    if (visited.has(name)) return;
    visited.add(name);
    const deps = getDeps(defs[name]);
    for (const dep of deps) {
      if (defs[dep]) visit(dep);
    }
    result.push(name);
  }

  for (const name of Object.keys(defs)) {
    visit(name);
  }

  return result;
}

function generateAllDefs(jsonSchema, mainName) {
  const defs = jsonSchema.$defs || jsonSchema.definitions || {};
  const output = [];
  const defOrder = topologicalSort(defs);

  for (const defName of defOrder) {
    const def = defs[defName];
    if (!def) continue;

    try {
      const resolvedDef = resolveRefs(def, defs);
      const zodCode = jsonSchemaToZod(resolvedDef, {
        module: 'none',
        name: undefined,
        withJsdoc: true,
      });

      output.push(`export const ${defName}Schema = ${zodCode};`);
      output.push(`export type ${defName} = z.infer<typeof ${defName}Schema>;`);
      output.push('');
    } catch (err) {
      console.error(`Error generating ${defName}:`, err.message);
      process.exit(1);
    }
  }

  // The root schema (invopop/jsonschema's ExpandedStruct: true inlines the
  // main type's own properties at the document root rather than as a
  // named $def), so it needs its own generation step keyed off mainName.
  try {
    const resolvedRoot = resolveRefs(jsonSchema, defs);
    const zodCode = jsonSchemaToZod(resolvedRoot, {
      module: 'none',
      name: undefined,
      withJsdoc: true,
    });

    output.push(`export const ${mainName}Schema = ${zodCode};`);
    output.push(`export type ${mainName} = z.infer<typeof ${mainName}Schema>;`);
    output.push('');
  } catch (err) {
    console.error(`Error generating root schema ${mainName}:`, err.message);
    process.exit(1);
  }

  return output.join('\n');
}

function main() {
  mkdirSync(OUTPUT_DIR, { recursive: true });

  for (const { file, name } of SCHEMAS) {
    const schemaPath = join(SCHEMA_DIR, file);
    console.log(`Processing ${file}...`);

    const jsonSchema = JSON.parse(readFileSync(schemaPath, 'utf-8'));
    const zodCode = generateAllDefs(jsonSchema, name);

    const outputPath = join(OUTPUT_DIR, `${name.toLowerCase()}.ts`);
    const fileContent = `/**
 * Auto-generated Zod schemas from JSON Schema.
 * DO NOT EDIT - regenerate with: npm run generate
 * Source: ${file}
 */

import { z } from 'zod';

${zodCode}
`;

    writeFileSync(outputPath, fileContent);
    console.log(`  -> Generated ${outputPath}`);
  }

  const indexContent = SCHEMAS.map(({ name }) => `export * from './${name.toLowerCase()}.js';`).join(
    '\n'
  );

  writeFileSync(
    join(OUTPUT_DIR, 'index.ts'),
    `/**
 * Auto-generated Zod schemas index.
 * DO NOT EDIT - regenerate with: npm run generate
 */

${indexContent}
`
  );

  console.log('Done!');
}

main();
