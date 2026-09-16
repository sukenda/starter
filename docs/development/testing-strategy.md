# Testing strategy

Prefer tests at the cheapest boundary that proves the behavior.

- Backend: unit tests for application logic, HTTP tests for transport contracts, MariaDB integration tests for SQL/repository behavior.
- Frontend: unit tests for schemas/helpers/stores and component tests for meaningful interactions; avoid snapshots as the primary assertion.
- Mobile: unit tests for state/domain behavior and widget tests for important user flows; platform integrations get focused adapter tests.

Do not mock your own code merely to increase coverage. Tests should protect contracts, invariants, authorization, error handling, and regressions.
