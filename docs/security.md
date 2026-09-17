# Security Model

## Trust boundary

The backend is authoritative for identity and authorization. Frontend/mobile checks exist for UX and must never be treated as a security control.

## Passwords

Passwords are hashed using Argon2id. Plaintext passwords exist only at the credential-verification boundary and must never be logged or persisted.

## Opaque session credentials

Access and refresh tokens are cryptographically random opaque values. Persist hashes, not raw token secrets. Access tokens are short-lived relative to refresh tokens. Refresh credentials rotate and single-use behavior is protected with a conditional persistence update.

## Browser session

Browser access tokens are held only in application memory. Refresh tokens are HttpOnly cookies, `Secure` outside development, and scoped to the auth path where applicable. Refresh requests use double-submit CSRF protection. The readable CSRF cookie/value must be accessible to the SPA path while the refresh secret remains unreadable to JavaScript.

The default deployment assumption is same-origin/same-site frontend and API behavior compatible with the cookie configuration. A future cross-subdomain deployment must deliberately review cookie Domain, SameSite, CORS, Origin and CSRF behavior rather than assuming host-only cookies will work.

## Native session

Flutter receives access and refresh tokens in JSON and stores both through platform secure storage. Native refresh tokens are not browser cookies. Refresh rotation is serialized and rotated credentials replace old local values.

## Authorization

RBAC is many-to-many: users have roles and roles have permissions. Protected routes enforce explicit permissions server-side. Menus may reference permissions only to determine navigation visibility.

## HTTP protections

Shared HTTP middleware provides request IDs, access logging, structured error handling and security headers. Sensitive values such as passwords, bearer tokens, refresh tokens and CSRF secrets must not appear in logs.

## Known hardening still required

Production readiness work still includes rate limiting/brute-force protection, login timing mitigation for unknown users, absolute refresh/session lifetime, stronger refresh-family/reuse handling and cleanup, broader security/integration tests, trusted Origin policy review, dependency/container scanning, secrets/deployment controls and operational incident procedures.

These items are documented as gaps, not as implemented controls.
