-- Agent-Platform core schema (SurrealQL)
-- Follows Agent-Blueprint conventions: schema.org typing, W3C DIDs, A2A,
-- OPA policy decision logging, and graph edges for trust/delegation.
-- TODO: split into per-domain files as this grows.

-- Agents are identified by a W3C DID and carry a trust level.
DEFINE TABLE agent SCHEMAFULL;
DEFINE FIELD did            ON agent TYPE string ASSERT $value != NONE;       -- W3C DID
DEFINE FIELD name           ON agent TYPE string;
DEFINE FIELD card           ON agent TYPE object;                            -- A2A AgentCard (see agent.proto)
DEFINE FIELD trust_level    ON agent TYPE int    DEFAULT 0 ASSERT $value >= 0 AND $value <= 5;
DEFINE FIELD runtime_id     ON agent TYPE option<string>;                    -- -> RuntimeSpec.id
DEFINE INDEX agent_did      ON agent FIELDS did UNIQUE;

-- Actions track the lifecycle of work (schema:Action).
DEFINE TABLE action SCHEMAFULL;
DEFINE FIELD type           ON action TYPE string;                           -- schema:Action subtype
DEFINE FIELD agent          ON action TYPE record<agent>;
DEFINE FIELD status         ON action TYPE string
  ASSERT $value IN ['potential', 'active', 'completed', 'failed'];
DEFINE FIELD started_at     ON action TYPE datetime DEFAULT time::now();
DEFINE FIELD ended_at       ON action TYPE option<datetime>;

-- Append-only event audit trail / telemetry.
DEFINE TABLE event SCHEMAFULL;
DEFINE FIELD kind           ON event TYPE string;
DEFINE FIELD subject        ON event TYPE option<record>;
DEFINE FIELD data           ON event TYPE object;
DEFINE FIELD at             ON event TYPE datetime DEFAULT time::now();

-- A2A messages exchanged between agents.
DEFINE TABLE a2a_message SCHEMAFULL;
DEFINE FIELD from_agent     ON a2a_message TYPE record<agent>;
DEFINE FIELD to_agent       ON a2a_message TYPE record<agent>;
DEFINE FIELD task_id        ON a2a_message TYPE string;
DEFINE FIELD payload        ON a2a_message TYPE object;
DEFINE FIELD at             ON a2a_message TYPE datetime DEFAULT time::now();

-- Skill / capability registry (A2A AgentCard skills).
DEFINE TABLE skill SCHEMAFULL;
DEFINE FIELD id_            ON skill TYPE string;
DEFINE FIELD name           ON skill TYPE string;
DEFINE FIELD tags           ON skill TYPE array<string>;

-- OPA policy decision log: every enforced decision is recorded.
DEFINE TABLE policy_decision SCHEMAFULL;
DEFINE FIELD agent          ON policy_decision TYPE record<agent>;
DEFINE FIELD policy         ON policy_decision TYPE string;
DEFINE FIELD allow          ON policy_decision TYPE bool;
DEFINE FIELD reason         ON policy_decision TYPE option<string>;
DEFINE FIELD at             ON policy_decision TYPE datetime DEFAULT time::now();

-- Graph edges: trust and delegation between agents.
DEFINE TABLE trusts   SCHEMAFULL TYPE RELATION FROM agent TO agent;
DEFINE TABLE delegates SCHEMAFULL TYPE RELATION FROM agent TO agent;
