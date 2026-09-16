.PHONY: setup infra-up infra-down dev dev-backend dev-frontend dev-mobile fmt lint test check

setup:
	cd apps/backend && go mod tidy
	cd apps/frontend && pnpm install
	cd apps/mobile && flutter pub get

infra-up:
	docker compose up -d mariadb

infra-down:
	docker compose down

dev:
	@echo "Start infrastructure first with: make infra-up"
	@echo "Then run applications in separate terminals:"
	@echo "  make dev-backend"
	@echo "  make dev-frontend"
	@echo "  make dev-mobile"

dev-backend:
	cd apps/backend && go run ./cmd/api

dev-frontend:
	cd apps/frontend && pnpm dev

dev-mobile:
	cd apps/mobile && flutter run

fmt:
	cd apps/backend && gofmt -w $$(find . -name '*.go' -not -path './vendor/*')
	cd apps/frontend && pnpm format
	cd apps/mobile && dart format lib test

lint:
	cd apps/backend && go vet ./...
	cd apps/frontend && pnpm typecheck && pnpm lint
	cd apps/mobile && flutter analyze

test:
	cd apps/backend && go test -race ./...
	cd apps/frontend && pnpm test
	cd apps/mobile && flutter test

check: lint test
