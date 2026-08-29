import { readFileSync, readdirSync } from 'fs';
import { join, resolve } from 'path';
import { describe, expect, it } from 'vitest';

import { ArchitectureSchema } from '../generated/architecture.js';

// Validates the same shared fixture corpus the Go side validates
// (schema/fixtures_test.go), so Go and Zod are proven to accept and
// reject the same documents rather than merely generating from the same
// schema.
const VALID_DIR = resolve(__dirname, '../../../examples/fixtures/valid');
const INVALID_DIR = resolve(__dirname, '../../../examples/fixtures/invalid');
const DOGFOOD_DIR = resolve(__dirname, '../../../examples/dogfood');

describe('ArchitectureSchema valid fixtures', () => {
  const files = readdirSync(VALID_DIR).filter((f) => f.endsWith('.json'));

  it('finds at least one fixture', () => {
    expect(files.length).toBeGreaterThan(0);
  });

  for (const file of files) {
    it(`validates ${file}`, () => {
      const raw = readFileSync(join(VALID_DIR, file), 'utf-8');
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

describe('ArchitectureSchema invalid fixtures', () => {
  const files = readdirSync(INVALID_DIR).filter((f) => f.endsWith('.json'));

  it('finds at least one fixture', () => {
    expect(files.length).toBeGreaterThan(0);
  });

  for (const file of files) {
    it(`rejects ${file}`, () => {
      const raw = readFileSync(join(INVALID_DIR, file), 'utf-8');
      const doc = JSON.parse(raw);
      const result = ArchitectureSchema.safeParse(doc);
      expect(result.success).toBe(false);
    });
  }
});

describe('ArchitectureSchema dogfood architectures', () => {
  const files = readdirSync(DOGFOOD_DIR).filter((f) => f.endsWith('.json'));

  it('finds at least one dogfood architecture', () => {
    expect(files.length).toBeGreaterThan(0);
  });

  for (const file of files) {
    it(`validates ${file}`, () => {
      const raw = readFileSync(join(DOGFOOD_DIR, file), 'utf-8');
      const doc = JSON.parse(raw);
      const result = ArchitectureSchema.safeParse(doc);
      if (!result.success) {
        throw new Error(`${file} failed validation: ${result.error.message}`);
      }
      expect(result.data.nodes.length).toBeGreaterThan(0);
      expect(result.data.relationships?.length ?? 0).toBeGreaterThan(0);
    });
  }
});
