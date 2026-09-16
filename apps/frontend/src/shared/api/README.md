# Frontend API boundary

Feature API adapters call `apiRequest` rather than using `fetch` directly from Vue components. The transport boundary owns base URL, headers, envelope parsing, and normalized infrastructure errors. Feature adapters own endpoint paths and feature-specific request/response schemas.
