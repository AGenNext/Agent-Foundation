# src/

Service code for Agent-Platform.

- **TypeScript** for the browser-side layer (per Agent-Platform's stack).
- **SurrealQL** holds business logic and policy enforcement close to the data.
- **AgentQL** (Langium DSL) compiles agent definitions to SurrealDB artifacts.

TODO: implement the source-to-artifact workflow:
`ingest → evidence → synthesis → decision → artifact/publication`, calling OPA
for enforcement and writing to the schema in `../database/schemas/`.
