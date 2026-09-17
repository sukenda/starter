# Backend Architecture

## Stack and entrypoint

The backend lives in `apps/backend`, is written in Go, exposes a versioned HTTP API, and uses MariaDB/MySQL through `database/sql`. `cmd/api/main.go` is the composition root: load/validate configuration, open the database, construct repositories/services/handlers, register middleware/routes, and start the server.

## Internal modules

```text
cmd/api/             composition root
internal/config/     environment loading and validation
internal/database/   DB opening/pool readiness helpers
internal/httpx/      HTTP concerns: envelopes, request/access logging, headers
internal/auth/       password hashing, tokens, sessions, auth HTTP/service/repository
internal/admin/      users, roles, permissions, menus HTTP/service/repository
migrations/          ordered SQL schema evolution
```

## Dependency rule

HTTP handlers translate protocol input/output. Services own use-case rules and depend on narrow store interfaces. Repositories implement persistence. New domain behavior should not place SQL in handlers or Fiber/HTTP concepts in repositories.

```text
HTTP route
  -> middleware
  -> handler
  -> service/use case
  -> Store interface
  -> SQL repository
  -> MariaDB
```

## HTTP boundary

API routes use `/api/v1`. Errors and success responses use shared envelope helpers. Request IDs and access logging provide correlation. Security headers are applied centrally. Health/readiness endpoints are infrastructure endpoints rather than domain use cases.

## Authentication

Passwords use Argon2id. Access and refresh credentials are opaque random tokens; only token hashes are persisted. Access tokens authenticate protected requests. Refresh tokens rotate through a compare-and-swap style persistence update so the same stored refresh credential cannot be successfully rotated twice.

The auth service owns login, refresh, current-user/session resolution, and logout. Browser cookie/CSRF behavior is an HTTP transport concern layered around the same session service. Native clients use JSON refresh tokens.

## Authorization / RBAC

Users can have multiple roles. Roles have multiple permissions. Authorization is permission-based at backend routes. The effective permission set is derived from user-role-role-permission relationships. Menus can reference a permission requirement but menus are navigation metadata, not the authorization boundary.

## Admin domain

`internal/admin` manages users, roles, permissions, and menus. Public API IDs are strings; database implementation details such as internal numeric IDs and SQL nullable types must not leak through DTOs. Multi-table writes such as role permission assignment are transactional. Menu parent relationships are resolved using public IDs and cycle/self-parent constraints are enforced.

## Persistence

Schema changes are forward migrations under `apps/backend/migrations`. Existing migrations establish users, roles, permissions, role assignments, sessions, menus, menu-permission integrity and related indexes/constraints. Application code must not silently mutate schema at runtime.

## Configuration

Configuration is loaded centrally and validated before serving traffic. Required database/auth/server values should fail fast. Auth access and refresh TTLs are explicit configuration and refresh TTL must exceed access TTL.

## Testing and quality

Backend CI runs formatting/vetting/tests including race detection. Unit tests currently cover selected config/database/http/auth primitives. Critical auth/admin integration coverage is still being expanded and must not be assumed complete.

## Extension pattern

For a new backend feature, create a domain-focused internal package when appropriate, define request/response DTOs at the HTTP edge, keep use-case rules in a service, define a narrow store interface beside the service, implement it in the repository, add migrations for persistence changes, protect routes with explicit permissions, update OpenAPI, tests, and docs in the same change.
