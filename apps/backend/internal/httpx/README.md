# HTTP transport foundation

`httpx` is deliberately limited to Fiber-aware transport primitives: API version grouping, envelopes, public errors, validation violations, request IDs, health endpoints, security headers, access logging, and centralized error mapping.

Business/domain packages must not import Fiber. Feature handlers translate HTTP input into application inputs and translate application results back through this package.
