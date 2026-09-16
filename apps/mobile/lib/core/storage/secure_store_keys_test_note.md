# Credential keys

Only long-lived credentials that require persistence should receive secure-store keys. Short-lived access tokens should remain in memory when the authentication design permits it. The auth PR will finalize token lifecycle after the backend contract is implemented.
