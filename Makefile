.PHONY: run
run:
	docker-compose up

.PHONY: run-build
run-build:
	docker-compose up --build