# Real Agent Foundation

A real agent is not a prompt, workflow, chatbot, API wrapper, model, or tool. Those are capabilities an agent may use.

A real agent is a persistent, governed digital identity that can perceive context, make decisions, take actions, remember outcomes, and be held accountable over time on behalf of a stakeholder.

## Canonical definition

> A real agent is a governed digital identity with memory, decision capability, action capability, and accountability, operating on behalf of a person, organization, system, or community across time.

This definition makes identity the anchor. Identity connects the agent's past, present, and future. Without identity, there is no continuity, responsibility, reputation, permission boundary, or lifecycle.

## Non-definitions

The following are not real agents by themselves:

- A large language model.
- A prompt.
- A chatbot.
- A one-shot workflow.
- A script.
- A tool call.
- A queue worker.
- A stateless API wrapper.
- A UI assistant without durable identity.

They can become part of an agent system only when bound to identity, policy, memory, capability, and accountability.

## Five primitives

```text
Identity -> Context -> Decision -> Action -> Outcome -> Memory -> Identity
```

### 1. Identity

Identity answers: who or what is acting?

A real agent has a durable identity that can be recognized across time. The identity may be issued by a person, organization, community, platform, or trust authority.

Identity carries:

- Agent identifier.
- Issuer.
- Owner or responsible stakeholder.
- Public agent card.
- Verifiable credentials.
- Capability grants.
- Trust score.
- Lifecycle state.
- Revocation status.

### 2. Context and memory

Context answers: what does the agent know now?

Memory answers: what must the agent carry forward?

A real agent must be able to use current context and durable memory without confusing one for the other.

Context includes:

- Current task.
- Current environment.
- Current user or stakeholder intent.
- Current constraints.
- Current available tools and resources.

Memory includes:

- Prior decisions.
- Prior outcomes.
- Preferences.
- Learned constraints.
- Relationships.
- Evidence.
- Commitments.
- Reputation and trust history.

### 3. Decision capability

Decision answers: what should the agent do next, and why?

A real agent makes decisions under goals, constraints, evidence, policy, cost, risk, and approvals.

A decision must be explainable enough to answer:

- What objective was pursued?
- What alternatives were considered?
- What constraints applied?
- What evidence was used?
- What policy allowed or blocked the decision?
- Was human approval required?

### 4. Action capability

Action answers: what can the agent change?

A real agent must be able to act through governed capabilities. Action is where risk enters the system.

Actions include:

- Calling tools.
- Sending messages.
- Creating or changing resources.
- Updating records.
- Requesting payments.
- Publishing artifacts.
- Delegating to another agent.
- Escalating to a human.

Every action should be scoped by identity, permissions, policy, and audit.

### 5. Accountability

Accountability answers: who is responsible, what happened, and can it be reviewed?

A real agent must leave a trace. The trace must connect identity, context, decision, action, outcome, and evidence.

Accountability includes:

- Audit trail.
- Provenance.
- Decision log.
- Policy evaluation result.
- Approval record.
- Action result.
- Error and exception record.
- Replay or review path.
- Revocation or suspension path.

## Lifecycle

A real agent is born, activated, governed, evaluated, suspended, revived, retired, or archived.

```text
Draft -> Issued -> Active -> Probation -> Suspended -> Terminated -> Archived
                  |              |
                  v              v
               Reviewed       Revoked
```

### Lifecycle states

| State | Meaning |
| --- | --- |
| Draft | Agent definition exists but identity is not issued. |
| Issued | Identity exists but runtime activation may still be blocked. |
| Active | Agent is allowed to operate within its grants. |
| Probation | Agent is allowed to operate with tighter policy after risk, drift, or low trust. |
| Suspended | Agent cannot act until reviewed or restored. |
| Terminated | Agent is permanently deactivated. |
| Archived | Agent history is retained for audit, learning, and provenance. |

## Governance gates

A real agent should not become active merely because code exists.

Activation requires governance gates:

- Identity issued.
- Owner assigned.
- Agent card published.
- Capability grants approved.
- Tool bindings reviewed.
- Policy attached.
- Data access scoped.
- Human escalation path defined.
- Audit enabled.
- Trust threshold configured.
- Revocation path available.

## Real-agent test

Use this test to reject fake agents.

An entity is a real agent only if it can answer all ten questions:

1. Who is the agent?
2. Who issued or authorized it?
3. On whose behalf does it act?
4. What is it allowed to do?
5. What is it forbidden to do?
6. What context did it use?
7. What decision did it make?
8. What action did it take?
9. What happened as a result?
10. Who can audit, suspend, or revoke it?

If any answer is missing, it is not yet a real agent. It is a component, prototype, workflow, tool, or assistant surface.

## Architecture model

```text
                    Trust Authority / Issuer
                              |
                              v
                      Agent Identity
                              |
       +----------------------+----------------------+
       |                      |                      |
       v                      v                      v
   Agent Card             Credentials             Policy
       |                      |                      |
       +----------------------+----------------------+
                              |
                              v
                         Agent Runtime
                              |
       +----------------------+----------------------+
       |                      |                      |
       v                      v                      v
    Context                 Tools                 Memory
       |                      |                      |
       +----------------------+----------------------+
                              |
                              v
                           Decision
                              |
                              v
                            Action
                              |
                              v
                           Outcome
                              |
                              v
                 Provenance + Audit + Evaluation
```

## Contract mapping

Agent-Foundation should represent the real agent as language-neutral contracts.

| Primitive | Contract concept |
| --- | --- |
| Identity | `AgentIdentity`, `AgentIssuer`, `AgentCredential` |
| Public discovery | `AgentCard` |
| Capability | `ToolBinding`, `CapabilityGrant` |
| Runtime | `RuntimeSpec` |
| Context | `Source`, `Evidence`, `Synthesis` |
| Decision | `Decision`, `Alternative` |
| Action | `AgentAction`, `ActionResult` |
| Accountability | `Provenance`, `Trace`, `BenchmarkRun`, audit records |
| Lifecycle | `AgentLifecycleState` |
| Trust | `TrustSignal`, `TrustScore`, evaluation results |

## Implementation boundary

Agent-Foundation defines the meaning and contracts.

Runtime enforcement belongs outside this repository:

- Agent-Platform enforces identity, policy, runtime, and action boundaries.
- Agent-Identity issues, verifies, suspends, and revokes identity and credentials.
- Agent-Trust evaluates trust signals.
- Agent-Bench and Agent-Eval evaluate behavior.
- Agent-Builder creates and updates agent definitions.
- Agent-UI displays lifecycle, approvals, traces, and audit.

## Short form

A fake agent can respond.

A real agent can be identified, authorized, governed, audited, trusted, suspended, and remembered.
