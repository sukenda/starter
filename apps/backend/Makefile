MIGRATE_VERSION ?= v4.19.1
MIGRATIONS_DIR ?= migrations
DATABASE_URL ?= starter:starter@tcp(127.0.0.1:3306)/starter?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci

.PHONY: tidy fmt vet test check migrate-up migrate-down

tidy:
	go mod tidy

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

vet:
	go vet ./...

test:
	go test -race ./...

check: tidy vet test

migrate-up:
	go run -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION) -path $(MIGRATIONS_DIR) -database "mysql://$(DATABASE_URL)" up

migrate-down:
	go run -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION) -path $(MIGRATIONS_DIR) -database "mysql://$(DATABASE_URL)" down 1
