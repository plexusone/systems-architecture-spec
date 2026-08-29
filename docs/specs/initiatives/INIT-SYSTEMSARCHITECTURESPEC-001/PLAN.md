# PLAN — Systems Architecture Spec — Machine-Readable System Contract (SAS v0.1)

**Initiative:** `INIT-SYSTEMSARCHITECTURESPEC-001`

## Sequencing Rationale

The sequencing follows the prove-value-first principle learned from the existing spec family: structured-changelog and structured-evaluation earned adoption through constant internal use, while specs without daily use are still proving themselves. SAS therefore optimizes for the first real user — our own PlexusOne portfolio (`plexusone/{agentforge,actionforge,dashforge}`) — and its two nearest payoffs: **diagrams** (replace hand-drawn architecture docs across repos that already share consistent specs/CHANGELOG.json structure) and **threat modeling for live-site launches** (product apps going to production on the web).

Phase 1 locks the core IR and the Go → JSON Schema → Zod pipeline because the IR is the hardest part to retrofit. Phase 2 proves optional-in-core/required-by-profile on real portfolio architectures. Phase 3 delivers payoff #1 (diagrams). Phase 4 delivers payoff #2 and closes v0.1: PIDL `ProtocolBinding` (PIDL is our own proven format, already used in threat-model-spec and aistandardsio/agent-protocols), the Threat Model Spec bridge, launch-readiness validation, and a real launch dogfooded end to end. Semantic diff moves to Phase 5 (v0.2 candidate) — the IR is designed so diff is expressible, but the engine ships after the launch payoffs. CALM interop is out of scope entirely: CALM is not yet widely adopted and is FINOS/finance-domain-centric, so adapters get scoped as new RMIs only if someone asks for CALM — it never blocks initiative completion.

## Phases

1. **Core IR & Schema Pipeline** — Architecture/Node/Relationship/Boundary/Identity/Entitlement/Extensions Go types; generated schemas linted by schemakit; embedded; round-trip fixtures. Exit: a hand-authored real architecture parses, validates structurally, and round-trips Go ↔ JSON ↔ Zod.
2. **Views, Profiles & Validation** — view-as-query, profile-conditional validation rules, `sas validate`, fixture corpus including one horizontal forge app and one vertical web app modeled end to end. Exit: security profile correctly flags a boundary-crossing edge lacking identity/encryption; development profile accepts the same file.
3. **Rendering & Catalogs** — Mermaid/D2/DOT renderers driven by views; AWS/GCP/K8s technology catalogs. Exit: three visually distinct, correct projections (context, deployment, security) of a real portfolio system render from `sas view`, good enough to replace that repo's hand-drawn diagrams.
4. **Threat Modeling, PIDL & Launch Readiness (closes v0.1)** — PIDL `ProtocolBinding` + `sas bind`; Threat Model Spec bridge exporting the system-under-analysis; launch-readiness rules for internet-facing exposure; assurance references + `sas assure`; a real live-site launch modeled and threat-modeled end to end. Exit: v0.1 tagged with SPEC.md, README, CI; a production launch used SAS artifacts for its threat model and readiness checks.
5. **Semantic Diff & Change Impact (v0.2 candidate)** — typed ChangeSet, ChangeImpact classification, baselines/assessments, FedRAMP change-assessment profile, `sas diff` CI output. Exit: the four canonical scenarios (new external dependency, authn change, read→read+write, new boundary crossing) classify correctly.

CALM interop is intentionally not a phase: it is demand-driven work scoped only when a real consumer asks.

## Milestones

- **M1 (end Phase 1):** first real architecture file round-trips through the full pipeline; schemas lint clean.
- **M2 (end Phase 3):** SAS replaces hand-drawn diagrams for at least one horizontal and one vertical portfolio app — one model, three rendered views each.
- **M3 (end Phase 4, v0.1):** a live web launch ships with a SAS-derived threat model and passing launch-readiness validation; v0.1 tagged.
- **M4 (end Phase 5, v0.2):** `sas diff` produces a defensible change-impact report on a real architecture change.

## Risks & Mitigations

- **Core ontology bloat (UML/ArchiMate trap).** Every proposed core field must pass the four-question test (diagrams / security analysis / change impact / reconciliation); everything else goes to extensions or catalogs. Kind taxonomy stays dotted-vocabulary small.
- **Authoring burden kills adoption.** We are the first user, so the feedback loop is immediate: portfolio examples act as authoring UX tests — if modeling a real PlexusOne app feels painful, the schema changes, not the example. Profiles keep the minimal file minimal.
- **Diagram quality below hand-drawn standard.** M2 explicitly requires "good enough to replace" for real repos, not synthetic demos; D2 is the primary target because it has the best generated-diagram ergonomics, with layout hints via view `groupBy` and catalog display data.
- **Threat-model bridge produces noise instead of signal.** The bridge exports only what Threat Model Spec defines as system-under-analysis (nodes, crossings, data flows, identities); threat reasoning stays in threat-model-spec — validated on a real launch in Phase 4, not hypothetically.
- **Diff deferral lets the IR drift away from diff-expressibility.** The expressiveness test (goal 6 in the PRD) is applied during Phase 1 design reviews even though the engine ships in Phase 5.
- **Zod/TS drift from Go model.** Shared fixture corpus is the conformance contract; CI runs all validators against it.

## Conventions

- Commits carry `Refs: RMI-SYSTEMSARCHITECTURESPEC-NNN` trailers; phases are reviewed and executed as units; phase status derives from member RMIs.
- Go-first schema workflow per org conventions: invopop/jsonschema generation, `schemakit lint --property-case camelCase`, `//go:embed`, `tools.go` guard for generator deps.
- Pre-push: `go test ./...`, `golangci-lint run`, coverage badge per global checklist.
