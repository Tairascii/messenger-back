BIN_DIR=${CURDIR}/bin
PACKAGE=messenger/chats

bindir:
	mkdir -p ${BIN_DIR}

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1 && \
	GOBIN=$(BIN_DIR) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0

.PHONY: create-migration
create-migration:
	$(BIN_DIR)/goose -dir db/migrations create -s $(name) sql

.PHONY: lint
lint:
	$(BIN_DIR)/golangci-lint run

.PHONY: run
run:
	export CONFIG_FILE=${CURDIR}/envs/values_local.yaml && \
	go run ${CURDIR}/cmd/chats/main.go
