# Systems Architecture Spec (SAS)

A statically-typed-friendly, Go-first specification for describing software systems as
semantic graphs: nodes, typed relationships, boundaries, identities, entitlements, and
protocol bindings.

**Architecture is a typed graph with properties and constraints; diagrams are views over
that graph.**

## Primary responsibility

The primary responsibility of this project is the semantic model itself — a spec whose
typed graph carries enough meaning that machines can reason over an architecture without
human interpretation. Diagrams, threat modeling, launch readiness, and change analysis
are **use-case requirements**: they determine what semantics the core must express and
prove the model works in practice, but they are consumers of the spec, not the spec.

If a use case needs something the model can't express, the model changes. If only one
use case wants something, it belongs in a consumer, a profile, or a namespaced extension
— never in the core.

## Status

Early development. See [`docs/specs/initiatives/INIT-SYSTEMSARCHITECTURESPEC-001/`](docs/specs/initiatives/INIT-SYSTEMSARCHITECTURESPEC-001/)
for the PRD, TRD, PLAN, and ROADMAP.

## Design principles

- **Statically Typed Friendly.** Go structs are the source of truth. JSON Schema is
  generated (never hand-written) and stays within the PlexusOne static-type profile:
  no `oneOf` / `anyOf` / `allOf`, no schema composition. Explicit discriminated fields
  instead of unions. TypeScript/Zod conform downstream. JSON is the interoperability
  contract — it never becomes a second type system.
- **Views are queries, not files.** One canonical model produces development,
  deployment, security, and threat-model projections without semantic drift.
  Renderers (Mermaid, D2, Graphviz DOT) are adapters, never sources of truth.
- **Profiles, not levels.** Semantics are optional in the core and required only by
  the profile that needs them, so a minimal architecture file stays minimal.
- **Core ontology, not a technology catalog.** "AWS Lambda" is not a core type — it's
  `kind: compute.function` + `technology: {provider: aws, service: lambda}`. Concrete
  vendor semantics live in catalogs outside the core schema.
- **Integration layer, not universal schema.** SAS binds to specialized specs by
  reference — [PIDL](https://github.com/grokify/pidl) for protocol choreography,
  [Threat Model Spec](https://github.com/grokify/threat-model-spec) for risk reasoning,
  [Multi-Agent Spec](https://github.com/plexusone/multi-agent-spec) for agent
  definitions — rather than absorbing them.

## License

MIT — see [LICENSE](LICENSE).
