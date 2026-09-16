# AI Engineering Contract

This file defines repository-wide rules for human and AI contributors.

## Before changing code

1. Read `ARCHITECTURE.md`.
2. Read the nearest `AGENTS.md` for the application being changed.
3. Inspect `packages/api-contract/openapi.yaml` for API-related work.
4. Prefer existing patterns over introducing a new abstraction or dependency.

## Repository boundaries

- `apps/backend`: Go API and backend business logic.
- `apps/frontend`: Vue web application.
- `apps/mobile`: Flutter mobile application.
- `packages/api-contract`: canonical HTTP API contract.
- `packages/design-tokens`: cross-client visual tokens; no framework-specific components.
- `docs`: architecture and engineering documentation.

Do not import source code directly across `apps/*`. Cross-application communication happens through explicit contracts.

## Feature workflow

When implementing a feature:

1. Identify affected applications.
2. Define/update the API contract first when HTTP behavior changes.
3. Implement backend behavior and tests.
4. Implement affected clients against the contract.
5. Add or update tests at the appropriate level.
6. Update architecture/documentation if a convention or system behavior changed.

## Dependency policy

- Do not add dependencies when the standard library or an existing dependency is sufficient.
- New dependencies must solve a concrete recurring problem.
- Prefer maintained, well-documented libraries with narrow responsibilities.
- Do not replace foundational libraries without documenting the decision in an ADR.

## General engineering rules

- Keep business logic out of transport/UI layers.
- Validate input at system boundaries.
- Return structured errors; do not expose internals to clients.
- Never commit secrets, credentials, tokens, `.env`, generated build output, or local IDE state.
- Add tests for business behavior and bug fixes.
- Keep functions/components small enough to have one clear responsibility.
- Prefer explicit code over clever abstractions.
- Preserve backward compatibility unless a breaking change is intentional and documented.

## API rules

- `/api/v1` is the initial public API namespace.
- `packages/api-contract/openapi.yaml` is the source of truth for public HTTP behavior.
- Use consistent error envelopes and HTTP status semantics.
- Pagination, filtering, sorting, authentication, and validation conventions must be consistent across endpoints.

## AI completion checklist

Before declaring work complete, verify:

- relevant formatters ran;
- lint passes;
- tests pass;
- API contract matches implementation;
- no secret or generated junk was added;
- documentation reflects architecture changes;
- unrelated files were not modified.
