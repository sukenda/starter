# HTTP security policy

The API applies conservative non-content security headers centrally. Browser-specific CSP belongs to the frontend delivery layer because the API does not render HTML. CORS will be configured from explicit environment allowlists; wildcard origins are permitted only for deliberate local-development scenarios.
