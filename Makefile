BIN_DIR=$(CURDIR)/bin
LOCAL_PG_PORT=6432
CHATS_MIGRATIONS_DIR=./chats/migrations
USERS_MIGRATIONS_DIR=./user/migrations
LOCAL_CHATS_DSN=postgres://postgres@localhost:$(LOCAL_PG_PORT)/chats-db?sslmode=disable
LOCAL_USERS_DSN=postgres://postgres@localhost:$(LOCAL_PG_PORT)/users-db?sslmode=disable

bindir:
	mkdir -p ${BIN_DIR}

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

.PHONY: run
run:
	docker-compose up

.PHONY: run-build
run-build:
	docker-compose up --build

.PHONY: migrations-up
migrations-up:
	$(BIN_DIR)/goose -dir $(CHATS_MIGRATIONS_DIR) postgres "$(LOCAL_CHATS_DSN)" up && \
	$(BIN_DIR)/goose -dir $(USERS_MIGRATIONS_DIR) postgres "$(LOCAL_USERS_DSN)" up
