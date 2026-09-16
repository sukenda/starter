# Engineering Foundation Quality Gate

This phase must be complete before authentication, RBAC, dashboard, user management, menu management, or role management are implemented.

## Backend

- Configuration is explicit, validated, and environment driven.
- MariaDB/MySQL uses a bounded connection pool and startup connectivity check.
- Database schema changes use versioned migrations; development data uses explicit seed commands.
- HTTP API has versioned routes, standardized success/error envelopes, request validation, request IDs, structured logging, recovery, CORS, and security headers.
- Liveness and readiness are separate concerns.
- Shutdown is graceful.
- Business features follow transport -> service/use-case -> repository boundaries without introducing interfaces that have no consumer/test seam.
- Tests cover core services and HTTP behavior; database integration tests run against MariaDB.

## Frontend

- Vue + TypeScript remains strict.
- Router, query/cache, state, HTTP transport, environment config, error normalization, and auth-session boundaries live outside presentation components.
- Server state uses TanStack Query; Pinia is reserved for client/session state.
- Feature modules own their pages, components, API adapters, schemas, and tests.
- Lint, typecheck, tests, and production build are mandatory CI gates.

## Mobile

- Riverpod owns dependency/state composition; GoRouter owns navigation; Dio owns HTTP transport.
- Secrets/tokens must use platform secure storage, never SharedPreferences.
- Network errors are normalized before reaching presentation.
- Features are organized by domain with data/domain/presentation separation where complexity warrants it.
- Formatting, static analysis, tests, and a build-level smoke check are CI gates.

## Cross-application

- `packages/api-contract/openapi.yaml` is the HTTP contract source of truth.
- API changes update the contract and consumers in the same stacked change when possible.
- No design-system implementation is added during this phase; only technical seams required to consume it later.
- No feature phase starts while a foundation CI gate is red.

## Planned stacked pull requests

1. `feat/production-monorepo-starter` -> `main`
2. `feat/engineering-foundation` -> `feat/production-monorepo-starter`
3. `feat/auth-rbac-foundation` -> `feat/engineering-foundation`
4. `feat/frontend-admin-shell` -> `feat/auth-rbac-foundation`
5. `feat/mobile-auth-home` -> latest accepted feature branch
6. design-system work starts only after the application milestones are stable.
