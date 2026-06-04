# Agent-Foundation

**Open community education and awareness initiative for the agent ecosystem.**

Agent-Foundation is the canonical non-profit and non-commercial foundation initiative of AGenNext. It supports community welfare through agent education, public awareness, research, open standards, and shared ecosystem contracts.

This repository defines language-neutral contracts for the AGenNext agent ecosystem. These contracts describe sources, evidence, decisions, evaluations, artifacts, theories, publications, agents, provenance, and runtime specifications in a portable format.

Agent-Foundation does not implement commercial products, runtime behavior, storage systems, or user-facing applications. Those responsibilities belong in implementation and product repositories such as Agent-Platform, Agent-Builder, Agent-UI, Agent-Research, Agent-Bench, and related AGenNext repositories.

## Design principles

- **Community welfare first.** The foundation exists to support responsible public understanding, education, and adoption of agentic systems.
- **Non-profit and non-commercial.** This initiative supports education, awareness, research, standards, and interoperability. It does not own product delivery, customer deployments, billing, or monetized implementation logic.
- **No database lock-in.** These are pure message contracts. There is no schema tied to SurrealDB, Postgres, or any other store. Persistence is the responsibility of each consuming service; Agent-Foundation only describes the shape of the data on the wire.
- **Language-neutral.** Contracts are defined once in `.proto`. This repo ships the `.proto` source only — each consuming repo generates its own bindings in Python, TypeScript, Go, Rust, Java, or any other language supported by `buf` or `protoc`.
- **Traceable by construction.** Evidence references sources, decisions reference evidence, artifacts and publications reference both. Nothing is asserted without provenance — _No evidence without source. No strategic decision without evidence._
- **Standards-aligned.** Contracts should interoperate with established open standards wherever possible instead of creating unnecessary proprietary concepts.
- **Stable by default.** Breaking changes must move to a new version package. Existing versioned contracts should remain compatible for downstream consumers.

## Scope

Agent-Foundation is responsible for:

- Community welfare through open agent education and awareness.
- Shared agent ecosystem terminology.
- Open, language-neutral contracts.
- Research-grounded data models.
- Standards alignment.
- Public-good documentation.
- Cross-repository interoperability.

Agent-Foundation is not responsible for:

- Commercial product delivery.
- Runtime enforcement.
- SaaS implementation.
- User interface development.
- Billing, sales, or monetization.
- Customer-specific deployments.
- Proprietary implementation logic.

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

Agent-Foundation only defines these contracts. Agent-Platform implements and enforces them at runtime; this repo carries no runtime or storage code.

## Mental model

Three planes, one contract package:

```text
            CONTROL PLANE
            Agent -> AgentCard -> ModelRouting -> ToolBinding -> RuntimeSpec
                    |
                    v
            KNOWLEDGE PLANE
            Source -> Evidence -> Synthesis -> Decision
                    |
                    v
            OUTPUT PLANE
            Theory -> BenchmarkRun -> ClearScore -> Publication
```

- **Control plane** — what an agent is: its card, models, tools, and runtime.
- **Knowledge plane** — what an agent knows, and where that knowledge came from.
- **Output plane** — what an agent produces, and how that output was evaluated.

`Provenance` threads through all three so any output can be traced back to its sources.

## Standards alignment

These contracts are designed to interoperate with established standards rather than reinvent them.

**Encoded in the contracts:**

| Standard | Where |
| --- | --- |
| W3C PROV-O | `Provenance` (`wasDerivedFrom` / `wasGeneratedBy` / `wasAttributedTo`) |
| OpenTelemetry GenAI semantic conventions | `Trace.model` / `Trace.provider`, `CostMetric` token fields |
| A2A / Agent2Agent | `AgentCard`, `AgentSkill`, `AgentCapabilities` |
| MCP / OpenAPI | `ToolBinding.protocol` |
| CSL / DOI / ORCID | `Source.doi`, `Publication.doi`, `Author.orcid` |
| OpenAI-compatible multi-LLM APIs | `ModelRouting`, `ModelRef.endpoint` |
| Kubernetes / Knative / Dapr / WASM | `RuntimeSpec`, `RuntimeKind` |

**Referenced and enforced by runtime repositories, not by these contracts:**

NIST AI RMF · ISO/IEC 42001 · OWASP LLM Top 10 · MITRE ATLAS · SPIFFE/SPIRE · Open Policy Agent · CycloneDX AI-BOM · CloudEvents.

## Research loop

```text
Theory -> BenchmarkRun + ClearScore -> Publication
  ^                                      |
  |                                      v
  +------------- Source <- Evidence <- Synthesis
```

## Quick start

This repo is proto only. It uses `buf` to keep the contracts healthy; code generation happens in the consuming repos.

```bash
buf lint
buf breaking --against '.git#ref=HEAD~1'
```

Consumers generate bindings from the `.proto` source, or pull a published Buf module, in their own repos rather than vendoring generated code here.

## Versioning

Contracts are versioned by package suffix, such as `v1`. Breaking changes go into a new version package, such as `v2`, never into an existing one, so consumers can migrate on their own schedule.
