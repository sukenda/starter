# Health endpoints

- `GET /health/live` proves the process can serve HTTP. It does not query downstream services.
- `GET /health/ready` proves required dependencies are available and currently checks MariaDB/MySQL connectivity.

Orchestrators should use liveness for restart decisions and readiness for traffic admission. Expensive business checks do not belong in either endpoint.
