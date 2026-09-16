# Transactions

Transactions are application/use-case concerns. A use case that coordinates multiple repository writes owns the transaction boundary so atomicity is explicit. Repository methods should accept the narrow query/exec capability they need when transactional reuse is required; avoid global transaction state.
