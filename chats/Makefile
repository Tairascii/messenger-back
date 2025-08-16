BIN_DIR=$(CURDIR)/bin
PACKAGE=messenger/chats
MIGRATIONS_DIR=./migrations
LOCAL_PG_PORT=6432
LOCAL_PG_DSN=postgres://postgres@localhost:$(LOCAL_PG_PORT)/chats-db?sslmode=disable

bindir:
	mkdir -p ${BIN_DIR}

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1 && \
	GOBIN=$(BIN_DIR) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0 && \
	GOBIN=$(BIN_DIR) go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: migrations-up
migrations-up:
	$(BIN_DIR)/goose -dir $(MIGRATIONS_DIR) postgres "$(LOCAL_PG_DSN)" up

.PHONY: create-migration
create-migration:
	$(BIN_DIR)/goose -dir ./migrations create -s $(NAME) sql

.PHONY: lint
lint:
	$(BIN_DIR)/golangci-lint run

.PHONY: run
run:
	export CONFIG_FILE=$(CURDIR)/envs/values_local.yaml && \
	go run $(CURDIR)/cmd/chats/main.go

.PHONY: db-up
db-up:
	docker-compose up

.PHONY: env-up
env-up: db-up

.PHONY: generate-swagger
generate-swagger:
	$(BIN_DIR)/swag fmt
	$(BIN_DIR)/swag init -g /cmd/chats/main.go -o docs
