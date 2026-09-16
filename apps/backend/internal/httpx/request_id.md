# Request correlation

Clients may send `X-Request-ID`; otherwise the API generates a cryptographically random 128-bit identifier. The value is returned in the response and included in structured access logs and public API errors so incidents can be correlated without exposing internal stack traces.
