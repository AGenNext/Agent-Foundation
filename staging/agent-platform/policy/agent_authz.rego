package agennext.agent.authz

# OPA authorization policy for Agent-Platform.
# Enforces, at runtime, invariants the contracts only describe.
# TODO: expand with OWASP LLM Top 10 / MITRE ATLAS controls feeding the
# AssuranceMetric.

import rego.v1

default allow := false

# An agent may act only if it has sufficient trust for the action's risk tier.
allow if {
	input.action.risk_tier <= input.agent.trust_level
	grounded
	not blocked
}

# "No decision without evidence": a publish/decision action must cite grounding.
grounded if {
	input.action.type != "decision"
}

grounded if {
	input.action.type == "decision"
	count(input.action.evidence_ids) > 0
}

# Deny anything the threat model flags (placeholder for ATLAS/OWASP checks).
blocked if {
	some flag in input.action.threat_flags
	flag in {"prompt_injection", "exfiltration", "policy_bypass"}
}

# Reason surfaced to the policy_decision audit log.
reason := "insufficient trust for risk tier" if {
	input.action.risk_tier > input.agent.trust_level
}

reason := "decision not grounded in evidence" if {
	input.action.type == "decision"
	count(input.action.evidence_ids) == 0
}
