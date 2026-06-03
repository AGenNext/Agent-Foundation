# Agent-Foundation

Shared, language-neutral contracts for the AGenNext agent ecosystem.

Agent-Foundation is the **base layer** every other AGenNext repo builds on. It
defines the core data contracts — sources, evidence, decisions, evaluations,
artifacts, theories, and publications — as [Protocol Buffers](https://protobuf.dev/),
so they can be shared across services written in any language.

## Design principles

- **No database lock-in.** These are pure message contracts. There is no schema
  tied to SurrealDB, Postgres, or any other store. Persistence is the
  responsibility of each consuming service; Agent-Foundation only describes the
  shape of the data on the wire.
- **Language-neutral.** Contracts are defined once in `.proto`. This repo ships
  the `.proto` source only — each consuming repo generates its own bindings
  (Python, TypeScript, or anything else `buf`/`protoc` supports).
- **Traceable by construction.** Evidence references sources, decisions
  reference evidence, artifacts and publications reference both. Nothing is
  asserted without provenance — _"No evidence without source. No strategic
  decision without evidence."_

## Contracts

| File | Contracts | Used by |
| --- | --- | --- |
| `source.proto` | `Source`, `Freshness`, `SourceType` | Agent-Research |
| `evidence.proto` | `Evidence`, `Synthesis`, `Confidence` | Agent-Research |
| `decision.proto` | `Decision`, `Alternative` | Agent-Research, Agent-Platform |
| `evaluation.proto` | `ClearScore` + CLEAR metrics, `Trace`, `BenchmarkRun` | Agent-Bench / CLEARBench |
| `artifact.proto` | `Artifact`, `ArtifactType` | Agent-Platform |
| `theory.proto` | `Theory`, `TheoryStatus` | Agent-Theories |
| `publication.proto` | `Publication`, `Author`, `PublicationType` | Agent-Publications |
| `provenance.proto` | `Provenance` (W3C PROV-O) | all entities |
| `agent.proto` | `Agent`, `AgentCard`, `ModelRouting`, `ToolBinding` | Agent-Platform |
| `runtime.proto` | `RuntimeSpec`, `RuntimeKind`, `ResourceRequirements` | Agent-Platform / Agent-Cloud |

All messages live in the `agennext.foundation.v1` package.

Agent-Foundation only **defines** these contracts. **Agent-Platform implements
and enforces** them at runtime; this repo carries no runtime or storage code.

## Mental model

Three planes, one contract package:

```
            ┌─────────────────────── CONTROL PLANE ───────────────────────┐
            │  Agent ── AgentCard (A2A) ── ModelRouting (multi-LLM)        │
            │    │           discover/call        primary + fallbacks      │
            │    └── ToolBinding (MCP / OpenAPI / A2A)                      │
            │    └── RuntimeSpec (k8s · serverless · wasm · dapr · edge)    │
            └──────────────────────────────────────────────────────────────┘
                                       │ runs
            ┌─────────────────────── KNOWLEDGE PLANE ─────────────────────┐
            │  Source ◄─ Retrieval(RAG) ◄─ Evidence ◄─ Synthesis ◄─ Decision │
            │     every step carries Provenance (PROV-O) ── full lineage   │
            └──────────────────────────────────────────────────────────────┘
                                       │ produces
            ┌─────────────────────── OUTPUT PLANE ────────────────────────┐
            │  Theory ─test▶ BenchmarkRun + ClearScore ─publish▶ Publication │
            │            (CLEAR: Cost·Latency·Efficacy·Assurance·Reliab.)   │
            └──────────────────────────────────────────────────────────────┘
```

- **Control plane** — what an agent *is*: its card, models, tools, runtime.
- **Knowledge plane** — what an agent *knows*, and where it came from.
- **Output plane** — what an agent *produces*, and how it was evaluated.

`Provenance` threads through all three so any output traces back to its sources.

## Standards alignment

These contracts are designed to interoperate with established standards rather
than reinvent them.

**Encoded in the contracts:**

| Standard | Where |
| --- | --- |
| **W3C PROV-O** (provenance) | `Provenance` (`wasDerivedFrom` / `wasGeneratedBy` / `wasAttributedTo`) |
| **OpenTelemetry GenAI** semconv | `Trace.model` / `Trace.provider`, `CostMetric` token fields |
| **A2A** (Agent2Agent) | `AgentCard`, `AgentSkill`, `AgentCapabilities` |
| **MCP** / **OpenAPI** | `ToolBinding.protocol` |
| **CSL / DOI / ORCID** | `Source.doi`, `Publication.doi`, `Author.orcid` |
| **OpenAI-compatible** multi-LLM | `ModelRouting`, `ModelRef.endpoint` |
| Kubernetes / Knative / Dapr / WASM | `RuntimeSpec`, `RuntimeKind` |

**Referenced, enforced by the runtime (Agent-Platform), not by these contracts:**
NIST AI RMF · ISO/IEC 42001 · OWASP LLM Top 10 & MITRE ATLAS (graded into the
CLEAR Assurance dimension) · SPIFFE/SPIRE (agent identity) · Open Policy Agent
(governance) · CycloneDX AI-BOM (supply chain) · CloudEvents (event envelope).

### The research loop, as contracts

```
Theory ──load──► BenchmarkRun + ClearScore ──test──► Publication ──publish──►
   ▲                                                      │
   └────────────── Source ◄─ Evidence ◄─ Synthesis ◄──────┘
```

## Quick start

This repo is **proto only**. It uses [`buf`](https://buf.build/docs/installation)
to keep the contracts healthy; code generation happens in the consuming repos.

```bash
# Lint the contracts
buf lint

# Check for breaking changes against the previous commit
buf breaking --against '.git#ref=HEAD~1'
```

Consumers generate bindings from the `.proto` source (or pull a published Buf
module) in their own repos, rather than vendoring generated code here.

## Versioning

Contracts are versioned by package suffix (`v1`). Breaking changes go into a new
version package (`v2`), never into an existing one, so consumers can migrate on
their own schedule.
