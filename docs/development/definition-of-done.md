# Definition of done

A change is done only when code, tests, contract/documentation, and CI agree.

For feature work:

1. Public API behavior is represented in OpenAPI.
2. Backend authorization and invariants are tested.
3. Client adapters normalize errors before presentation.
4. Loading, empty, error, and success states are intentional.
5. No credentials or sensitive payloads are logged.
6. Native formatter/linter/analyzer, tests, and build checks pass.
7. Architecture/ADR documentation is updated when a durable decision changes.

Passing locally is not sufficient when the pull-request quality gates are red.
