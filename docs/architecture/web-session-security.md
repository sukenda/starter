# Web and Mobile Session Security

## Credential transport

The starter deliberately uses different refresh-token transport for browser and native clients.

### Browser

- Access token is short-lived and kept only in application memory.
- Refresh token is stored in a `HttpOnly` cookie scoped to `/api/v1/auth`.
- The non-HttpOnly CSRF cookie is scoped to `/` so the SPA can read it from routes such as `/dashboard` and copy it into `X-CSRF-Token`.
- Both cookies are `Secure` outside the development environment and use `SameSite=Lax`.
- JavaScript never receives the browser refresh token.
- Refresh uses double-submit CSRF: the readable CSRF cookie must match `X-CSRF-Token`.
- Refresh requests use `credentials: include`.
- A browser reload restores the session by refreshing first and then calling `/auth/me`.
- Concurrent refreshes in the SPA are serialized through one in-flight promise.
- Cookie clearing must use the same Path attribute used when each cookie was created.

### Mobile

- Login and refresh continue to exchange refresh tokens in JSON.
- The native client owns persistence in platform secure storage.
- Browser cookies are not required for the mobile flow.

## Rotation

Refresh tokens are opaque, hashed at rest, rotated on every refresh, and updated with compare-and-swap semantics. A previously rotated token cannot rotate the session again.

## CSRF and XSS

`HttpOnly` protects the browser refresh credential from direct JavaScript access but does not remove the need to prevent XSS. The double-submit token protects the cookie-backed refresh endpoint against cross-site request forgery. State-changing endpoints authenticated with bearer access tokens do not rely on ambient cookies for authorization.

The CSRF cookie is intentionally readable and broadly path-scoped; it is not an authentication secret. The refresh token remains HttpOnly and narrowly scoped.

## Deployment assumptions

Production must terminate TLS before the API. The current cookie model assumes a browser/API topology where the SPA can receive/read the CSRF cookie and send the refresh cookie to the API. If frontend and API are hosted on different sites or hosts, review cookie Domain, `SameSite`, CORS credentials, trusted origins, and CSRF policy together; do not independently loosen one control.

## Future hardening

Products with stronger security requirements may add refresh-token families/reuse detection, session/device management, absolute session lifetime, login throttling, trusted-origin enforcement, and step-up authentication. These should be driven by the product threat model rather than added as generic complexity by default.
