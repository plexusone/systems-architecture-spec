# TRD — Systems Architecture Spec — Machine-Readable System Contract (SAS v0.1)

**Initiative:** `INIT-SYSTEMSARCHITECTURESPEC-001`

## Architecture Overview

```text
Go structs (source of truth)
      │ invopop/jsonschema
      ▼
JSON Schema  ──  schemakit lint --property-case camelCase  ──  //go:embed
      │
      ▼
Canonical JSON (interchange contract)
      │
      ├──► Zod schemas + TypeScript types (conformance target)
      ├──► Renderers: Mermaid / D2 / Graphviz DOT
      ├──► Semantic diff → Change Impact IR → compliance profiles
      └──► CALM import/export adapters
```

JSON Schema validates the wire contract; it must never become a second type system. The generated schema stays within the PlexusOne static-type profile: no `oneOf`/`anyOf`/`allOf`/`not`/conditional schemas, homogeneous arrays, deterministic property types, explicit discriminated fields for polymorphism.

## Core Graph Model

```go
type Architecture struct {
    Version       string            `json:"version"`
    Metadata      Metadata          `json:"metadata,omitempty"`
    Nodes         []Node            `json:"nodes"`
    Relationships []Relationship    `json:"relationships"`
    Boundaries    []Boundary        `json:"boundaries,omitempty"`
    Bindings      []ProtocolBinding `json:"bindings,omitempty"`
    Views         []View            `json:"views,omitempty"`
}
```

### Node

```go
type Node struct {
    ID          string       `json:"id"`
    Kind        NodeKind     `json:"kind"` // e.g. "service", "gateway", "compute.function", "data.database", "agent", "mcp_server"
    Name        string       `json:"name"`
    Owner       string       `json:"owner,omitempty"`
    Technology  *Technology  `json:"technology,omitempty"` // {provider: "aws", service: "lambda"}
    Identity    *Identity    `json:"identity,omitempty"`
    Boundaries  []string     `json:"boundaries,omitempty"` // IDs; multi-membership, no single hierarchy
    Criticality string       `json:"criticality,omitempty"` // tier0, tier1, ...
    Assurance   *Assurance   `json:"assurance,omitempty"`
    Refs        []ExternalRef `json:"refs,omitempty"`      // OpenAPI, Multi-Agent Spec, Threat Model Spec, IaC
    Extensions  Extensions    `json:"extensions,omitempty"` // namespaced typed maps (security, sre, agent, compliance)
}
```

Kind taxonomy is a dotted vocabulary (`compute.function`, `compute.instance`, `data.database`, `messaging.queue`, `gateway`, `service`, `agent`, `mcp_server`, `external_service`, `actor`), deliberately small; concrete products come from `Technology` + catalogs, never new core kinds.

### Relationship

The edge is the security-critical object. It must be able to express what most diagram formats collapse into an arrow:

```go
type Relationship struct {
    ID            string         `json:"id"`
    From          string         `json:"from"`
    To            string         `json:"to"`
    Kind          RelationKind   `json:"kind"` // calls | data.access | data_flow | publishes_to | subscribes_to | authenticates_to | authorizes_via | observes | enforces | depends_on | spawns | delegates_to
    Transport     *Transport     `json:"transport,omitempty"`     // protocol, port, applicationProtocol, encryption
    Operations    []Operation    `json:"operations,omitempty"`    // generic verb vocabulary
    Identity      *IdentityRef   `json:"identity,omitempty"`      // which identity the caller acts as
    Authorization *Authorization `json:"authorization,omitempty"` // entitlements
    Data          *DataFlow      `json:"data,omitempty"`          // classifications, e.g. customer_data, source_code
    CrossesBoundaries []string   `json:"crossesBoundaries,omitempty"` // derived + assertable
    CriticalPath  bool           `json:"criticalPath,omitempty"`
    Sync          string         `json:"sync,omitempty"`          // sync | async
    Assurance     *Assurance     `json:"assurance,omitempty"`
    ProtocolRef   string         `json:"protocolRef,omitempty"`   // e.g. "pidl://oauth/dpop-authorization-code"
    Extensions    Extensions     `json:"extensions,omitempty"`
}
```

**Operations vocabulary:** generic verbs `discover | read | create | update | delete | execute | administer` with catalog-supplied specialization mappings (HTTP GET → read; SQL SELECT → read; MCP invoke_tool → execute; Agent spawn/delegate → execute). Views choose the abstraction level; the model stores both when the author supplies specifics.

### Boundary

```go
type Boundary struct {
    ID         string     `json:"id"`
    Kind       string     `json:"kind"`         // network | trust | environment | account | region | compliance | organization
    Name       string     `json:"name"`
    Attributes map[string]string `json:"attributes,omitempty"` // e.g. vpc: prod, classification: restricted
    Compliance *ComplianceBoundary `json:"compliance,omitempty"` // e.g. fedrampBoundary: true, controls: ["SC-7", ...]
    Extensions Extensions `json:"extensions,omitempty"`
}
```

Nodes reference boundaries by ID (many-to-many). "Crossing" is computable: an edge crosses boundary B when exactly one endpoint is a member.

### Identity & Entitlement

First-class, not buried in attributes:

```go
type Identity struct {
    Type      string `json:"type"` // human | workload | service | agent | device | external
    Mechanism string `json:"mechanism,omitempty"` // aws_iam_role | spiffe | oauth_client | oidc | aauth | k8s_service_account | x509 | api_key
    Ref       string `json:"ref,omitempty"`
}

type Entitlement struct {
    Subject    string   `json:"subject"`
    Action     string   `json:"action"`
    Resource   string   `json:"resource"`
    Conditions []string `json:"conditions,omitempty"`
}
```

This maps naturally onto IAM systems and OpenFGA/SpiceDB tuples later, without making ReBAC part of the core.

### ProtocolBinding (PIDL integration)

PIDL models protocols with abstract participants; SAS instantiates them:

```go
type ProtocolBinding struct {
    ID          string            `json:"id"`
    ProtocolRef string            `json:"protocolRef"` // "pidl://oauth/dpop-authorization-code"
    Participants map[string]string `json:"participants"` // pidl participant → SAS node ID
}
```

The same PIDL OAuth/AAuth/MCP definition binds into many architectures. SAS never duplicates choreography; a lightweight `Flow` concept is deliberately omitted from v0.1 core in favor of `ProtocolBinding` + `protocolRef` on edges.

### External References & Extensions

```go
type ExternalRef struct {
    Type string `json:"type"` // openapi | multi-agent-spec | threat-model-spec | pidl | terraform | otel
    Ref  string `json:"ref"`
}
```

`Extensions` is a map of namespace → typed extension object (`security`, `sre`, `compliance`, `agent`). Extension shapes are declared as Go types per namespace — explicit fields, not schema composition. This is the SAS answer to CALM's `allOf`-based extension model.

### Assurance (evidence references only in v0.1)

```go
type Assurance struct {
    Tests      []string `json:"tests,omitempty"`      // test://integration/orders-db-authz
    Metrics    []string `json:"metrics,omitempty"`    // otel://service/orders-api
    Detections []string `json:"detections,omitempty"` // siem://rule/DB-UNEXPECTED-WRITE
    Deployment []string `json:"deployment,omitempty"` // aws://lambda/orders-api
}
```

SAS stores pointers to evidence, never the evidence itself. Coverage reporting (which declared elements lack tests/metrics/detections) is computable from the graph; reconciliation against live evidence is out of scope for v0.1.

## Views & Profiles

**A view is a query, not a diagram file:**

```go
type View struct {
    ID               string   `json:"id"`
    Name             string   `json:"name"`
    IncludeKinds     []string `json:"includeKinds,omitempty"`
    IncludeRelations []string `json:"includeRelations,omitempty"`
    IncludeBoundaries []string `json:"includeBoundaries,omitempty"`
    GroupBy          string   `json:"groupBy,omitempty"` // boundary kind to nest by
    Level            string   `json:"level,omitempty"`   // generic | specific (operations abstraction)
}
```

**A profile declares which optional semantics become mandatory** (validation, not schema variation):

| Profile | Requires |
|---|---|
| development | nodes, relationships |
| deployment | technology, boundaries (account/region/network), transport |
| security | identity on boundary-crossing edges, encryption, operations, entitlements on writes |
| threat-model | data classifications, boundary membership, external flags |
| sre | criticality, owner, availability target on tier0/tier1 |
| fedramp | compliance boundary membership, control mappings on boundary components |

Profile checks are application-level validation rules over one schema — the schema itself never branches.

## Validation Rules Engine

Rule examples shipped in v0.1:

```text
tier0 node          → must have owner, availability target, assurance.metrics
boundary crossing   → must define transport.encryption and identity (security profile)
write operations    → must define authorization.entitlements (security profile)
external node edge  → must declare data classifications (threat-model profile)
internet-facing     → public endpoints must define authn, TLS, owner,
                      data classification (launch-readiness rules)
dangling references → node/boundary/binding IDs must resolve
```

Rules are data-driven where practical (rule ID, severity, profile, predicate) so profiles compose. Launch-readiness rules are the go-live gate for portfolio web apps: `sas validate --profile security` must pass before a vertical app ships to production.

## Threat Model Spec Bridge

SAS supplies the system-under-analysis that `grokify/threat-model-spec` needs, so threat models never re-describe the system:

```text
SAS architecture ──► sas export threat-model ──► system-under-analysis
                                                  ├── components (nodes)
                                                  ├── trust-boundary crossings
                                                  ├── data flows + classifications
                                                  ├── identities per edge
                                                  └── external dependencies
```

The bridge exports facts only; threats, mitigations, and risk reasoning stay in Threat Model Spec, which references SAS IDs (`sas_resource`, `sas_relationship`). Dogfooded on the first live-site launches (vertical portfolio apps).

## Semantic Diff & Change Impact (ships post-v0.1)

The IR must stay diff-expressible from day one (the PRD expressiveness test), but the engine ships in v0.2.

```text
architecture@baseline ──┐
                        ├──► sas diff ──► ChangeSet (typed) ──► ChangeImpact ──► profile classification
architecture@proposed ──┘
```

- **ChangeSet:** typed changes — NodeAdded/Removed/Modified, RelationshipAdded/..., BoundaryMembershipChanged, field-level deltas (authentication mtls→oauth, operations read→read+write, internal→external target).
- **ChangeImpact:** semantic classification — `NEW_EXTERNAL_DEPENDENCY`, `AUTHN_MECHANISM_CHANGED`, `ENTITLEMENT_EXPANSION`, `NEW_BOUNDARY_CROSSING`, `DATA_CLASSIFICATION_CHANGE`, `COMPUTE_MODEL_CHANGED`, plus affected boundaries and (via mappings) potentially affected controls.
- **Compliance profiles** consume ChangeImpact: the FedRAMP profile emits `NO_REVIEW_REQUIRED | CHANGE_ASSESSMENT_REQUIRED | POTENTIAL_SIGNIFICANT_CHANGE` with the triggering changes as evidence. The tool produces decision support; classification as an actual significant change is a recorded human decision (`Assessment{id, decision, reviewer, evidence}`).
- **Baselines:** an approved baseline is a named, immutable architecture version; releases reference `{baseline, proposed, assessment}`.

## Renderers

Renderer = pure function `(Architecture, View) → string`. v0.1 targets: Mermaid (GitHub-native), D2 (best generated-diagram ergonomics), Graphviz DOT (auto-layout for large graphs). Structurizr/C4 and draw.io are explicitly later. Catalogs provide display labels/icon hints (AWS Lambda icon for `aws/lambda`).

## CALM Interop (out of scope; demand-driven)

CALM (FINOS) can never be the foundation: its core uses `anyOf` (node-type-definition) and recommends `allOf` for extension, which violates the static-type profile. Adapters are not scheduled in this initiative — CALM adoption is still limited and finance-domain-centric, so `sas import calm` / `sas export calm` get scoped as new work only when a real consumer asks. If built: mapping is straightforward at the instance level (Node↔Node, Relationship↔Relationship, Boundary↔deployed-in/composition, Controls↔compliance extension); lossy conversions are reported, never silent; core never depends on CALM shapes.

## CLI

```text
# v0.1
sas validate <arch.json> [--profile security,threat-model]
sas view <arch.json> --view security --format d2|mermaid|dot
sas bind <arch.json> --pidl <protocol.json>     # verify bindings resolve
sas export threat-model <arch.json>              # system-under-analysis for threat-model-spec
sas assure <arch.json>                           # assurance-reference coverage report

# v0.2+
sas diff <base.json> <proposed.json> [--profile fedramp] [--format json]

# demand-driven (not scheduled)
sas import calm <calm.json> / sas export calm <arch.json>
```

Library-first: all capability in importable packages; `cmd/sas` is a thin Cobra adapter.

## Repository Layout

```text
systems-architecture-spec/
├── sas/            # core types: architecture, node, relationship, boundary, identity, view, binding
├── validate/       # rules engine + profiles
├── diff/           # semantic diff, change impact, assessments
├── render/
│   ├── mermaid/
│   ├── d2/
│   └── dot/
├── catalog/        # technology catalogs (aws, gcp, k8s) + operation mappings
├── bridge/
│   └── threatmodel/ # system-under-analysis export for threat-model-spec
├── schema/         # generated *.schema.json + //go:embed + gen/ (go:build ignore) + tools.go
├── ts/             # Zod schemas + TS types + shared fixtures runner
├── cmd/sas/        # cobra CLI
├── examples/       # real architectures (dogfood) + valid/invalid fixture corpus
└── docs/specs/initiatives/INIT-SYSTEMSARCHITECTURESPEC-001/
```

## Conformance & Testing

- Shared fixture corpus (`examples/fixtures/{valid,invalid}`) exercised by the Go validator, JSON Schema validation, and Zod — behavioral conformance, not just type generation.
- Round-trip test: Go → JSON → (schema validate) → Zod parse → JSON → Go; semantic equality required.
- `schemakit lint --property-case camelCase` in CI, **default profile**, matching established org practice (`deppolicy`). The `scale` profile additionally rejects any use of `additionalProperties` as a value schema — including the strict, homogeneous `map[string]string` on `Boundary.Attributes` — which is a stronger bar than the static-type-friendly principle requires: the principle bans dynamic, heterogeneous, ambiguous shapes (`map[string]any`, unions), not a fixed-value-type map that Go, Rust, and TypeScript (`Record<string,string>`) all represent identically and unambiguously.
- `tools.go` with `//go:build tools` blank imports so `go mod tidy` keeps the schema generator's deps (standard gen/main.go pattern).
- TypeScript/Zod generation follows the proven `multi-agent-spec` pipeline: `json-schema-to-zod` per `$def` (topologically ordered, refs inlined for self-contained output) plus one generation pass for the root document type, keyed by name (`ArchitectureSchema`) since `invopop/jsonschema`'s `ExpandedStruct: true` inlines the root type's own properties rather than emitting it as a named `$def`.

## Dependencies

Cobra (CLI), `invopop/jsonschema` (schema generation), `grokify/schemakit` (lint, dev/CI tool). Verify latest versions before adding. No runtime dependency on PIDL/Threat Model Spec/Multi-Agent Spec — references are by URI; optional typed helpers may import their types later.
