# API versioning

Public application endpoints live below `/api/v1`. Operational endpoints such as liveness/readiness remain outside the public API namespace. Breaking public-contract changes require an explicit versioning decision rather than silently changing an existing endpoint.
