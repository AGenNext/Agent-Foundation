# Agent-Foundation

**Open community education and awareness initiative for the agent ecosystem.**

Agent-Foundation is the canonical non-profit and non-commercial foundation initiative of AGenNext. It supports community welfare through agent education, public awareness, research, open standards, and shared ecosystem contracts.

This repository defines language-neutral contracts for the AGenNext agent ecosystem. These contracts describe sources, evidence, decisions, evaluations, artifacts, theories, publications, agents, provenance, and runtime specifications in a portable format.

Agent-Foundation does not implement commercial products, runtime behavior, storage systems, or user-facing applications. Those responsibilities belong in implementation and product repositories such as Agent-Platform, Agent-Builder, Agent-UI, Agent-Research, Agent-Bench, and related AGenNext repositories.

## Mission

To make agent education, awareness, research, standards, and interoperability openly accessible for the benefit of communities, learners, researchers, builders, institutions, and organizations.

Agentic systems are becoming part of how people learn, work, build, decide, govern, and collaborate. Communities need open knowledge, shared vocabulary, responsible guidance, and interoperable foundations so they can understand and adopt these systems safely.

Agent-Foundation exists to reduce confusion, fragmentation, and unequal access by publishing open educational material, shared terminology, research-grounded models, and portable contracts for the agent ecosystem.

## Scope

Agent-Foundation is responsible for:

- Community welfare through open agent education and awareness.
- Shared agent ecosystem terminology.
- Open, language-neutral contracts.
- Research-grounded data models.
- Standards alignment.
- Public-good documentation.
- Cross-repository interoperability.
- Provenance and traceability guidance.
- Evaluation and benchmarking literacy.
- Responsible adoption guidance.

Agent-Foundation is not responsible for:

- Commercial product delivery.
- Runtime enforcement.
- SaaS implementation.
- User interface development.
- Billing, sales, or monetization.
- Customer-specific deployments.
- Proprietary implementation logic.

## Principles

- **Community welfare first.** Foundation work should improve public understanding, responsible adoption, and shared access to agent knowledge.
- **Education before adoption.** People should be able to understand agentic systems before they are expected to use, govern, buy, or deploy them.
- **Evidence before claims.** Important claims should be supported by evidence, sources, rationale, or clearly stated assumptions.
- **Provenance by design.** Sources, evidence, decisions, evaluations, artifacts, theories, and publications should be traceable.
- **Open standards before proprietary definitions.** Agent-Foundation should align with open standards wherever possible.
- **Interoperability over lock-in.** Foundation contracts should work across languages, tools, agents, runtimes, and databases.
- **Stability over novelty.** Shared contracts should be stable enough for other repositories and communities to depend on.
- **Public good before commercial advantage.** Commercial product delivery, monetization, customer-specific implementation, and proprietary runtime behavior belong outside this repository.
- **Transparency over opacity.** Definitions, decisions, trade-offs, limitations, and assumptions should be documented clearly.
- **Responsible adoption over hype.** The foundation should support practical education, risk understanding, evaluation literacy, and community welfare.

## Governance

Agent-Foundation is authoritative for:

- Foundation terminology used across AGenNext repositories.
- Canonical language-neutral contracts.
- Shared public-good education and awareness material.
- Standards alignment notes.
- Versioning and compatibility expectations for foundation contracts.
- Public documentation explaining agent concepts, evidence, decisions, evaluation, provenance, and interoperability.

Agent-Foundation is not authoritative for:

- Commercial product roadmaps.
- Customer-specific implementations.
- Production runtime behavior.
- User interface design.
- Billing, packaging, sales, or monetization.
- Proprietary deployment architecture.

Changes should be reviewed against four questions:

1. Does this support community welfare, education, awareness, research, standards, or interoperability?
2. Is this a foundation contract or public-good artifact rather than product implementation?
3. Is the terminology stable enough to be reused across repositories?
4. Does the change preserve backward compatibility for existing versioned contracts?

Breaking changes must not be introduced into an existing version package. Use a new package version such as `v2` when compatibility cannot be preserved.

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
