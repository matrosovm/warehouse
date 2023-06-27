define db_env
PG_DSN=postgres://localhost:5432/db_warehouse?sslmode=disable&timezone=UTC
endef

LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run
run:
	go build -o bin/ ./... && ./bin/warehouse


bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/mitchellh/gox@v1.0.1  && \
	go install golang.org/x/tools/cmd/goimports@v0.1.9 && \
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0 \

bin-local-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pav5000/smartimports/cmd/smartimports@v0.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

.PHONY: format
format: bin-local-deps
	PATH=$(LOCAL_BIN):$(PATH) smartimports -exclude 'internal/pb,pkg/api'

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint

.PHONY: lint
lint: bin-local-deps
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...

db-status: $(eval $(call db_env))
	PATH=$(LOCAL_BIN):$(PATH) goose -dir migrations postgres "$(PG_DSN)" status

db-up: $(eval $(call db_env))
	@echo "up migrations"
	PATH=$(LOCAL_BIN):$(PATH) goose -dir migrations postgres "$(PG_DSN)" up

db-down: $(eval $(call db_env))
	@echo "down migration"
	PATH=$(LOCAL_BIN):$(PATH) goose -dir migrations postgres "$(PG_DSN)" down

# Для использования команды нужно прописать: make db-create NAME=%NAME_MIGRATION%
db-create: $(eval $(call db_env))
	@echo "create migration '$(NAME)'"
	PATH=$(LOCAL_BIN):$(PATH) goose -dir migrations postgres "$(PG_DSN)" create "$(NAME)" sql

db-dump: $(eval $(call db_env))
	pg_dump --schema-only "$(PG_DSN)" > schema.sql

pgdocker-run:
	docker run -d --name pglocal_warehouse -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${USER} -e POSTGRES_DB=db_warehouse -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgres:latest