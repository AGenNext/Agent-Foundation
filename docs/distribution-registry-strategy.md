# Distribution and Registry Strategy

Agent-Foundation should be distributable through trusted registries, package channels, and agent SDK ecosystems while preserving provenance, human accountability, and release gates.

## Core rule

```text
Publish everywhere useful.
Trust only what is gated.
```

## Distribution targets

```text
GitHub Releases
GitHub Container Registry / GHCR
Docker Hub
Zot OCI Registry
npm
Buf Registry
Anthropic Skills
OpenAI Agents SDK integrations
Downstream AGenNext repos
```

## Registry roles

| Registry / Channel | Role |
|---|---|
| GitHub Releases | canonical release records and provenance artifacts |
| GHCR | primary OCI/container image registry for GitHub-native distribution |
| Docker Hub | public container distribution and discovery |
| Zot Registry | self-hosted OCI registry for air-gapped, sovereign, edge, and internal deployments |
| npm | TypeScript SDKs, schema packages, MCP helpers, generated types |
| Buf Registry | protobuf module publication and contract consumption |
| Anthropic Skills | packaged contribution/review workflows |
| OpenAI Agents SDK | integration adapter path for agents consuming foundation contracts |

## Trust requirements

Every published artifact should include:

```text
source repository
source commit SHA
version
build workflow
workflow run ID
artifact digest
provenance manifest
human approval signal
release notes
```

## Package naming

Suggested names:

```text
GHCR: ghcr.io/agennext/agent-foundation
Docker Hub: agennext/agent-foundation
Zot: registry.<domain>/agennext/agent-foundation
npm: @agennext/foundation
npm: @agennext/foundation-mcp
npm: @agennext/foundation-schemas
Buf: buf.build/agennext/foundation
Anthropic Skills: agennext-foundation-contracts, agennext-foundation-review
OpenAI Agents SDK: agennext-foundation-agent-adapter
```

## Publication flow

```text
Commit
  -> CI
  -> human accountability gate
  -> trusted artifact build
  -> provenance manifest
  -> GitHub release
  -> registry publish
  -> downstream notification
```

## What can be published

```text
protobuf contract bundles
schema bundles
TypeScript generated types
MCP helpers
contract validation CLI
containerized validators
documentation bundles
Anthropic Skills for contract review
OpenAI Agents SDK adapters and examples
OCI images for validators and generators
```

## Zot registry profile

Zot is the recommended self-hosted OCI registry profile when deployments require:

```text
air-gapped operation
sovereign artifact custody
edge-local registry
internal enterprise distribution
low-footprint OCI registry
GitHub-independent fallback
```

Rule:

```text
Zot mirrors trusted artifacts. It does not create trust by itself.
```

## Agent SDK integrations

Anthropic Skills and OpenAI Agents SDK integrations are assistant-facing packaging layers.

They may help:

```text
review a contract proposal
check provenance completeness
prepare a vocabulary proposal
summarize compatibility impact
review release notes
check contribution against the foundation constitution
consume foundation contracts inside agent workflows
```

They must not become authority.

```text
Agent SDKs assist.
Human approves.
Foundation records truth.
```

## One-line definition

Distribution registry strategy means Agent-Foundation publishes trusted, provenance-backed artifacts to GitHub Releases, GHCR, Docker Hub, Zot, npm, Buf, Anthropic Skills, and OpenAI Agents SDK integration paths while keeping human accountability and release gates at the center.
