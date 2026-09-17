# API Guide

## Contract ownership

The backend implementation lives under `apps/backend`; the repository-level API contract lives under `packages/api-contract`. Endpoint, field, status-code or authentication changes must update both surfaces together.

## Base path

Application API endpoints are versioned under `/api/v1`. Infrastructure health/readiness routes may sit outside domain authorization according to the backend router.

## Response conventions

Success and error payloads use centralized HTTP envelope helpers. Error responses should expose stable application error codes/messages and request correlation where supported, not raw SQL/framework errors.

## Authentication endpoints

The implemented auth surface includes:

- `POST /api/v1/auth/login` — verifies credentials and creates a session. Browser and mobile transports differ after authentication.
- `POST /api/v1/auth/refresh` — rotates refresh credentials. Browser uses refresh cookie + CSRF; native uses JSON `refresh_token`.
- `GET /api/v1/auth/me` — returns the authenticated user/session representation including authorization context used by clients.
- `POST /api/v1/auth/logout` — invalidates the current session and clears transport credentials where applicable.

## Admin endpoints

The admin API exposes management resources for users, roles, permissions and menus under `/api/v1`. Routes are permission-protected. Current list APIs are baseline list operations; pagination/filter/sort is planned hardening and must not be assumed until implemented.

User writes support multi-role assignment. Role writes support permission assignment. Menu writes use public parent identifiers and optional permission metadata.

## Authentication header

Protected API calls use an access token as `Authorization: Bearer <token>`. Browser refresh tokens must not be sent as JavaScript-readable bearer credentials. Native refresh tokens are sent only to the refresh endpoint.

## CSRF

Cookie-authenticated browser refresh requires the double-submit CSRF value/header expected by the backend. CSRF is a browser transport concern; native JSON refresh does not rely on browser cookies.

## Authorization

Every protected backend route declares/enforces the required permission. Client route/menu permission checks are not a substitute.

## Compatibility

Within `/api/v1`, avoid accidental breaking changes. Prefer additive fields and explicit migration/deprecation for contract changes. When a breaking redesign is unavoidable, document it and consider a new API version rather than silently changing semantics.
