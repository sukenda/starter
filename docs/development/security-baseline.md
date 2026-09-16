# Security baseline

This starter treats client-side authorization as a user-experience concern only. The backend is the security boundary.

## Required practices

- Passwords are never logged or stored in plaintext.
- Authentication tokens and secrets are redacted from logs.
- Mobile refresh credentials use platform secure storage.
- Browser session design must account for XSS and CSRF before authentication is implemented.
- Authorization is enforced server-side for every protected operation.
- SQL values are parameterized; user input is never concatenated into SQL.
- Validation happens at the transport boundary and business invariants are enforced in application/domain code.
- Error responses expose stable public codes without leaking stack traces or infrastructure details.
- Request IDs allow support correlation without exposing sensitive context.
- CORS uses explicit allowed origins outside local development.
- Security headers and request size/time limits are configured centrally.
- Dependencies are pinned through ecosystem lock/checksum files and reviewed through CI.

Authentication and RBAC decisions are intentionally deferred to the next stacked PR so this foundation remains independently reviewable.
