# Charter

Agent-Foundation is the canonical non-profit and non-commercial foundation initiative of AGenNext.

It exists as an open community education and awareness initiative for the agent ecosystem, supporting community welfare through education, awareness, research, open standards, and shared ecosystem contracts.

## Purpose

The purpose of Agent-Foundation is to provide a stable public-good foundation for understanding, describing, evaluating, and interoperating with agentic systems.

Agent-Foundation helps communities, learners, researchers, builders, institutions, and organizations use a shared language for agent systems without depending on a single vendor, runtime, database, or commercial product.

## Authority

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

## Stewardship

Foundation stewardship should protect the following outcomes:

- Public usefulness.
- Community welfare.
- Responsible education.
- Clear terminology.
- Standards alignment.
- Contract stability.
- Interoperability across tools, agents, runtimes, and repositories.
- Separation between public-good foundation work and commercial implementation work.

## Scope of foundation work

Foundation work may include:

- Shared vocabulary.
- Protocol Buffer contracts.
- Data model documentation.
- Public education material.
- Awareness guides.
- Research summaries.
- Standards mapping.
- Compatibility policy.
- Provenance and traceability guidance.
- Evaluation and benchmarking concepts.

Foundation work should not include:

- Application source code.
- Runtime enforcement engines.
- SaaS product features.
- Customer deployment code.
- Billing systems.
- Sales workflows.
- Private customer data.
- Proprietary business logic.

## Relationship with AGenNext product repositories

AGenNext product and implementation repositories may consume Agent-Foundation contracts and terminology.

Examples include:

- Agent-Platform for runtime enforcement.
- Agent-Builder for agent creation workflows.
- Agent-UI for user-facing interfaces.
- Agent-Research for research and evidence workflows.
- Agent-Bench for evaluation and benchmarking.
- Agent-Registry for catalog and discovery.

Those repositories implement product behavior. Agent-Foundation defines shared public-good foundations.

## Change principles

Changes to Agent-Foundation should be accepted when they:

- Improve community welfare, education, awareness, research, standards, or interoperability.
- Clarify shared terminology.
- Strengthen contract stability.
- Improve traceability and provenance.
- Align with open standards.
- Preserve the separation between foundation and commercial implementation.

Changes should be rejected or moved elsewhere when they:

- Introduce product-specific behavior.
- Add customer-specific logic.
- Create vendor lock-in.
- Break existing contracts without a version migration path.
- Mix commercial delivery responsibilities into the foundation layer.

## Compatibility

Published foundation contracts should be stable within their version package.

Breaking changes must move to a new version package, such as `v2`. Existing consumers should be able to migrate intentionally instead of being forced into silent breaking changes.

## North star

Agent-Foundation should remain a public-good initiative that improves community welfare by making agent education, awareness, research, standards, and interoperability easier to understand, share, and adopt responsibly.
