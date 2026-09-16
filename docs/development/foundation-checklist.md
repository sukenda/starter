# Foundation completion checklist

A foundation pull request remains draft until all applicable boxes are complete.

## CI and reproducibility

- [ ] Backend dependency metadata is committed and `go mod tidy` is clean.
- [ ] Frontend pnpm lockfile is committed and install uses frozen lockfile.
- [ ] Flutter pub lockfile is committed.
- [ ] Backend vet/test pass.
- [ ] Frontend lint/typecheck/test/build pass.
- [ ] Mobile format/analyze/test pass.

## Backend runtime

- [x] Environment configuration boundary.
- [x] Configuration invariant validation.
- [x] MariaDB/MySQL connection pool.
- [x] Startup database connectivity check.
- [x] Versioned migration convention.
- [x] Structured JSON logging baseline.
- [x] Graceful shutdown.
- [x] Standard API envelope/error model.
- [x] Central error handler.
- [x] Separate liveness/readiness endpoints.
- [x] Versioned API group.
- [ ] Request ID and access logging middleware.
- [ ] Recovery, CORS, security headers, body/time limits.
- [ ] Request validation primitive.
- [ ] MariaDB integration-test job.

## Frontend runtime

- [x] Strict TypeScript baseline.
- [x] HTTP transport boundary.
- [x] Normalized API errors.
- [x] Environment configuration.
- [x] TanStack Query / Pinia responsibility documented.
- [x] Session boundary prepared without implementing auth UI.

## Mobile runtime

- [x] Compile-time environment configuration.
- [x] Dio transport boundary.
- [x] Normalized API errors.
- [x] Riverpod dependency boundary.
- [x] Secure-storage abstraction prepared without implementing auth UI.

Authentication/RBAC and visual design-system implementation are explicitly out of scope for this PR.
