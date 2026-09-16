# Secure storage

Authentication credentials must be persisted through the `SecureStore` abstraction. The authentication milestone will provide a platform implementation backed by OS-protected secure storage. Do not store access/refresh credentials in SharedPreferences, plain files, logs, or application state snapshots.
