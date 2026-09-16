# Dependency policy

Use standard-library/platform capabilities when they are clear and sufficient. Add a dependency when it removes meaningful complexity or provides a well-maintained ecosystem standard.

Dependencies must be pinned by the native ecosystem lock/checksum mechanism, used behind an application boundary when replacement cost is material, and kept out of domain logic where practical. Avoid adding packages for trivial helpers.
