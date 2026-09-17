# Operations Guide

## Current runtime shape

The backend is the network service and depends on MariaDB/MySQL. Vue is a browser build that calls the backend API. Flutter is a native client that calls the same API. Production container/deployment manifests are not yet a completed baseline and should not be inferred from local development files.

## Configuration

Backend environment configuration is centralized in `internal/config` and represented by `.env.example`. Validate configuration before accepting traffic. Secrets belong in the deployment secret mechanism, never committed environment files.

Frontend/mobile API endpoint configuration must point to the deployed backend origin appropriate for the environment. Browser cookie/CSRF behavior must be validated whenever frontend/API origins change.

## Database

Apply ordered SQL migrations as a controlled deployment step before application code that requires the new schema becomes active. Backups, restore verification and rollback/runbook automation are production-hardening work still to be added.

## Health

Use the backend health/liveness endpoint for process health and readiness for dependency-aware traffic admission. Database readiness uses a real DB readiness check. Do not put authenticated business behavior behind infrastructure probes.

## Logging

Backend requests receive request IDs and access logs. Preserve correlation IDs through incident investigation. Never log credentials, cookies, bearer/refresh tokens, passwords or secret configuration.

## CI

Backend, frontend and mobile have independent GitHub Actions workflows. A merge should only be treated as validated when the relevant workflow has completed successfully; a merge that occurs while CI is still running can still expose a post-merge failure.

## Deployment assumptions

The browser session implementation is safest with an explicitly designed same-origin/same-site topology. Cross-origin or cross-subdomain deployments require deliberate CORS/cookie/CSRF/Origin configuration review.

## Production gaps / roadmap

Before calling the starter a complete production deployment baseline, implement and verify reproducible lockfiles/toolchains, multi-stage non-root container images, graceful shutdown, metrics/tracing, rate limiting/security hardening, automated image/security scanning, release/versioning workflow, migration deployment strategy, backup/restore procedure, rollback procedure, secret management guidance, monitoring/alerting and incident runbooks.

Keep this list synchronized with implementation: remove an item only when the code/config/runbook actually exists and is tested.
