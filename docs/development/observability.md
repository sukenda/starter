# Observability baseline

Logs are structured and correlated by request ID. Operational endpoints distinguish liveness from readiness. Future metrics and traces should reuse the same service/request context and must not expose credentials or sensitive payloads.

Instrumentation belongs at infrastructure/transport boundaries; business logic should remain testable without an observability SDK.
