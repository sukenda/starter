# Configuration

Configuration is loaded once at process startup and validated before infrastructure is initialized. Environment variables are deployment inputs; feature code receives typed configuration rather than reading `os.Getenv` directly. Secrets must not have production defaults or be logged.
