# Database migrations

Migrations are ordered, immutable after deployment, and reversible when practical.

Rules:

- Never edit a migration that has reached a shared environment; add a new migration instead.
- Use InnoDB, `utf8mb4`, and explicit indexes/constraints.
- Keep schema migration separate from application startup in production.
- Seeds are not migrations. Production reference data must be introduced deliberately and idempotently.
- Destructive changes should use expand/migrate/contract when zero-downtime deployment matters.

Local commands are exposed through `apps/backend/Makefile`.
