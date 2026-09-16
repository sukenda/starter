# Authentication and RBAC

## Decisions

- Passwords are hashed with Argon2id. Password hashes are never reversible.
- Access and refresh credentials are cryptographically random opaque tokens.
- Only SHA-256 token hashes are persisted in MariaDB; plaintext tokens exist only at issuance and on the client.
- Access tokens are short lived (15 minutes by default).
- Refresh tokens are long lived (30 days by default) and rotated on every successful refresh.
- Logout revokes the server-side session immediately.
- Authentication is server-stateful by design so revocation works consistently across multiple API instances without per-instance token state.
- Authorization is permission based. Roles are collections of permissions; users may have multiple roles.
- Backend permission checks are the security boundary. Frontend/mobile permission checks are UX only.

## Data model

```text
users ──< user_roles >── roles ──< role_permissions >── permissions
  │
  └──< auth_sessions
```

`auth_sessions` stores hashes for both access and refresh tokens. A session has independent access and refresh expirations and an explicit revocation timestamp.

## Request flow

Login verifies an active user and Argon2id password hash, creates a persistent session, and returns access + refresh tokens. Protected requests authenticate the bearer token against the session store and resolve effective permissions. Refresh validates the refresh token and atomically replaces both credentials. Logout marks the session revoked.

## Operational notes

Expired/revoked sessions should be periodically deleted by an operational cleanup job once the application has a scheduler/background-worker strategy. Authentication failures intentionally use generic public errors to avoid account enumeration.
