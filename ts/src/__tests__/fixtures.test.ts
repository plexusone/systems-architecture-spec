import { readFileSync, readdirSync } from 'fs';
import { join, resolve } from 'path';
import { describe, expect, it } from 'vitest';

import { ArchitectureSchema } from '../generated/architecture.js';

// Validates the same shared fixture corpus the Go side validates
// (schema/fixtures_test.go), so Go and Zod are proven to accept the same
// documents rather than merely generating from the same schema.
const FIXTURES_DIR = resolve(__dirname, '../../../examples/fixtures/valid');

describe('ArchitectureSchema fixtures', () => {
  const files = readdirSync(FIXTURES_DIR).filter((f) => f.endsWith('.json'));

  it('finds at least one fixture', () => {
    expect(files.length).toBeGreaterThan(0);
  });

  for (const file of files) {
    it(`validates ${file}`, () => {
      const raw = readFileSync(join(FIXTURES_DIR, file), 'utf-8');
      const doc = JSON.parse(raw);
      const result = ArchitectureSchema.safeParse(doc);
      if (!result.success) {
        throw new Error(`${file} failed validation: ${result.error.message}`);
      }
      expect(result.data.version).toBeTruthy();
      expect(result.data.nodes.length).toBeGreaterThan(0);
    });
  }
});
