# Freeze, Do Not Destroy

Agent-Foundation uses a freeze-not-destroy emergency model.

When trust is uncertain, all running processes that can change trusted state must freeze. Nothing is destroyed. Evidence, history, issues, discussions, and recovery paths remain available.

## Core rule

```text
Freeze changing processes.
Preserve all evidence.
Destroy nothing.
Recover through human accountability.
```

## What freezes

```text
merge readiness
release promotion
trusted artifact publishing
contract promotion
workflow-based publishing
project state promotion
foundation truth publication
adapter/runtime promotion
```

## What remains open

```text
read access
issues
incident reports
discussions
forensic review
reconciliation issues
recovery branches
recovery pull requests
human review
```

## Why nothing is destroyed

Destruction removes evidence.
Evidence is required for accountability, reconciliation, and recovery.

Therefore:

```text
Do not delete history.
Do not erase artifacts.
Do not hide traces.
Do not rewrite evidence.
Do not destroy state.
```

## Frozen process model

```text
Normal process
  -> uncertainty detected
  -> kill switch active
  -> running trusted-state-changing processes freeze
  -> evidence preserved
  -> incident/reconciliation issue opened
  -> recovery PR created
  -> human accountability review
  -> gates reopened after reconciliation
```

## One-line definition

Freeze-not-destroy means the emergency control stops all trusted-state-changing processes while preserving evidence and allowing accountable human recovery.
