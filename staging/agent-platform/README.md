# Agent-Platform (scaffold)

> Staged in Agent-Foundation for hand-off. Move this folder to
> `AGenNext/Agent-Platform`.

Agent-Platform is the **assembly + enforcement layer** of the AGenNext
ecosystem. It composes the specialized subsystems into governed, source-to-
artifact intelligence workflows. It **implements and enforces** the contracts
defined in [Agent-Foundation](https://github.com/AGenNext/Agent-Foundation) and
follows the conventions of
[Agent-Blueprint](https://github.com/AGenNext/Agent-Blueprint) (SurrealDB graph,
schema.org typing, W3C DIDs, OPA policy, A2A messaging).

## What it does

1. Ingests **Sources** (incl. RAG retrievals) and builds **Evidence** →
   **Synthesis** → **Decision** with full **Provenance** (PROV-O).
2. Runs **Agents** (multi-LLM `ModelRouting`, A2A `AgentCard`, MCP tools) on a
   `RuntimeSpec` substrate via Agent-Runtime.
3. Evaluates systems with **CLEAR** (Cost, Latency, Efficacy, Assurance,
   Reliability) and emits **BenchmarkRun** reports.
4. Produces governed **Artifacts** and **Publications**, enforcing policy at
   each step (OPA) and recording an audit trail.

## Contract enforcement

This layer is where Agent-Foundation's contracts stop being advisory:

| Contract (Agent-Foundation) | Enforcement here |
| --- | --- |
| `Evidence.source_ids >= 1` | Reject ungrounded evidence at ingest |
| `Decision` grounded (CEL) | Block decisions lacking synthesis/evidence |
| `AssuranceMetric` | OPA policy + OWASP LLM Top 10 / MITRE ATLAS grading |
| `AgentCard` | Publish/validate A2A agent cards |
| `RuntimeSpec` | Hand to Agent-Runtime for scheduling |

## Layout

```
database/schemas/   SurrealQL schema (agents, actions, events, A2A, policy log)
policy/             OPA / Rego authorization policies
deploy/             open-container-compose.yml (local) + Helm (cluster)
sdk/                buf codegen config pulling Agent-Foundation contracts
src/                service code (TypeScript browser layer + SurrealQL logic)
```

## Quick start

```bash
# Generate the Agent-Foundation client SDK
buf generate --template sdk/buf.gen.yaml

# Bring up SurrealDB + platform services locally
docker compose -f deploy/open-container-compose.yml up
```

## Status

Skeleton only. Each directory contains a representative starting file with
`TODO`s; nothing here is production-ready.
