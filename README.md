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
| `publication.proto` | `Publication`, `PublicationType` | Agent-Publications |

All messages live in the `agennext.foundation.v1` package.

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
