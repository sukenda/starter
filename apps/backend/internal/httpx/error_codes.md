# Public API error codes

Error `code` values are machine-readable compatibility contracts. Human-readable `message` text may evolve; clients should branch on `code`, not message text.

Foundation codes:

- `internal_error`: unexpected server failure.
- `http_error`: generic HTTP-layer error.
- `validation_failed`: request shape or field validation failed.
- `service_unavailable`: a required dependency is not ready.

Feature PRs add domain-specific codes close to the owning feature and update the OpenAPI contract.
