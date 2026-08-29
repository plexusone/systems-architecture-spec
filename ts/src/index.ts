/**
 * Systems Architecture Spec (SAS) TypeScript SDK.
 *
 * The Go structs in package sas are the source of truth; these Zod
 * schemas and types are generated from the JSON Schema derived from them
 * (npm run generate). Do not hand-edit src/generated — regenerate it.
 */

export * from './generated/index.js';
