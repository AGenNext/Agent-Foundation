# Governance

Agent-Foundation is an open community education and awareness initiative for the agent ecosystem.

It is the canonical non-profit and non-commercial foundation initiative of AGenNext. Its purpose is to support community welfare through education, awareness, research, open standards, and shared ecosystem contracts.

## Governance principles

- **Community welfare first.** Foundation work should improve public understanding, safe adoption, responsible use, and shared access to agent knowledge.
- **Non-profit and non-commercial.** This repository does not own monetized product delivery, customer deployments, billing, sales, or proprietary implementation logic.
- **Open by default.** Foundation artifacts should be useful to learners, researchers, builders, institutions, and communities.
- **Standards-aligned.** Contracts and terminology should align with open standards wherever possible.
- **Traceable.** Important claims, models, contracts, and decisions should include provenance, source references, or rationale.
- **Stable.** Published contracts should remain compatible within a version. Breaking changes must move to a new version package.

## What belongs here

Agent-Foundation may contain:

- Shared agent ecosystem vocabulary.
- Language-neutral contracts.
- Public-good education and awareness material.
- Research-grounded models.
- Standards alignment notes.
- Community welfare resources.
- Cross-repository interoperability guidance.
- Versioning and compatibility rules.

## What does not belong here

Agent-Foundation should not contain:

- Commercial product implementation.
- Runtime enforcement code.
- SaaS application code.
- User interface implementation.
- Billing, sales, or monetization workflows.
- Customer-specific deployment logic.
- Proprietary implementation logic.

Commercial and implementation work belongs in product/runtime repositories such as Agent-Platform, Agent-Builder, Agent-UI, Agent-Research, Agent-Bench, and related AGenNext repositories.

## Decision model

Changes should be reviewed against four questions:

1. Does this support community welfare, education, awareness, research, standards, or interoperability?
2. Is this a foundation contract or public-good artifact rather than product implementation?
3. Is the terminology stable enough to be reused across repositories?
4. Does the change preserve backward compatibility for existing versioned contracts?

If the answer to any of these is unclear, the change should be documented before adoption.

## Contract changes

Contract changes should include:

- Problem statement.
- Proposed contract or terminology change.
- Compatibility impact.
- Standards alignment, where applicable.
- Migration notes, if needed.

Breaking changes must not be introduced into an existing version package. Use a new package version such as `v2` when compatibility cannot be preserved.
