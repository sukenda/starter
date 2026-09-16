# Backend AI Rules

Applies to `apps/backend`.

## Stack

- Go
- Fiber v2 for HTTP delivery
- PostgreSQL
- Prefer `pgx/v5` for database access when persistence is introduced

## Structure

- `cmd/api`: process entrypoint and dependency wiring.
- `internal/config`: environment/runtime configuration.
- `internal/platform`: reusable technical infrastructure.
- `internal/<feature>`: business feature code.

A typical feature may contain `handler.go`, `service.go`, `repository.go`, `model.go`, and `dto.go`, but create only the files/layers the feature actually needs.

## Rules

- Keep `main.go` focused on bootstrapping/wiring.
- Keep HTTP handlers thin: parse, validate, invoke behavior, translate result.
- Keep SQL/database details out of handlers.
- Pass `context.Context` through operations that perform I/O.
- Wrap errors with useful context; do not compare error strings.
- Use constructor functions for dependencies rather than package globals.
- Prefer interfaces at consumer boundaries when substitution/testing is actually needed; do not create an interface for every struct.
- Use table-driven tests where they improve clarity.
- Avoid ORM magic by default; explicit SQL/pgx is preferred for predictable data access.
- Never panic for expected runtime errors.
- Support graceful shutdown for long-running server processes.

## HTTP

- Keep public behavior aligned with `/packages/api-contract/openapi.yaml`.
- Use JSON consistently.
- Centralize error response translation.
- Do not expose stack traces/internal database errors.

## Validation

Validate transport input before invoking business behavior. Business invariants must still be enforced in the service/domain layer and must not rely solely on HTTP validation.
