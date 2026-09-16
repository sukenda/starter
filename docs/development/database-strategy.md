# MariaDB/MySQL strategy

MariaDB is the default local/runtime database target while the Go driver uses the MySQL protocol. SQL should stay within the common MariaDB/MySQL subset unless a project explicitly chooses a vendor-specific feature.

Use InnoDB, utf8mb4, explicit foreign keys/indexes, UTC timestamps at service boundaries, bounded pools, versioned migrations, parameterized SQL, and integration tests for repository queries. Schema migration is a deployment step rather than an implicit side effect of API startup.
