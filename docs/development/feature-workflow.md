# Feature Development Workflow

Use this workflow for features that can affect multiple applications.

## 1. Define behavior

Write down the user-visible behavior and identify whether the backend, web, mobile, database, or public API contract changes.

## 2. Update the contract

When HTTP behavior changes, update `packages/api-contract/openapi.yaml` before client/server implementation. Keep naming and error semantics consistent with existing endpoints.

## 3. Implement backend

Add the smallest feature structure that preserves boundaries. Keep transport parsing in handlers, business behavior in service/domain code, and persistence in repositories. Add tests around business behavior and boundary cases.

## 4. Implement clients

Frontend and mobile consume the same contract but retain native architecture. Do not copy server business rules into clients unless the rule is intentionally required for local UX; the server remains authoritative.

## 5. Verify states

For user-facing asynchronous workflows verify at least: loading, success/data, empty when applicable, validation failure, authorization failure, and unexpected/server failure.

## 6. Quality gate

Run:

```bash
make check
```

Then verify contract and documentation consistency.

## 7. Architecture changes

If the feature introduces a new cross-cutting architectural choice, create an ADR under `docs/decisions/` instead of silently establishing a new convention.
