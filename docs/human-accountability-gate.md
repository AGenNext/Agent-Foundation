# Human Accountability Gate

Agent-Foundation uses a human accountability gate for trusted merges, releases, and published artifacts.

Automation may validate, build, package, summarize, and prepare. A human remains accountable for approval.

## Core rule

```text
Automation prepares.
CI validates.
Human approves.
Maintainer merges.
Release publishes.
```

## Why this exists

GitHub artifacts are published content. Foundation contracts become downstream truth. A change to the foundation can affect platform, product, SDK, research, and partner work.

Therefore, trusted outputs require accountable human approval.

## Human accountability chain

```text
Contribution
  -> CI validation
  -> human review
  -> accountable approval
  -> merge
  -> artifact build
  -> provenance manifest
  -> release gate
  -> published artifact
```

## Trusted artifact requirements

A trusted artifact must include:

```text
source repository
source commit SHA
workflow run ID
build timestamp
artifact name
artifact digest
provenance manifest
human approval signal
release status
```

## Approval authority

The approval authority is the designated human maintainer or publishing authority for the repository.

For Agent-Foundation, AGenNext maintainers hold the human accountability gate.

## Non-negotiables

```text
No artifact without provenance.
No release without digest.
No publish without human approval.
No community contribution becomes foundation truth without review.
No automation may override the human gate.
```

## One-line definition

Human accountability gate means automation may accelerate work, but an accountable human authority must approve trusted merges, releases, and published artifacts before they become foundation truth.
