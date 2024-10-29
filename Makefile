SHELL=/bin/bash
export REPO=grpc

.PHONY: run-local
run_local:
	@go run cmd/main.go

.PHONY: up_db
up_db:
	@docker run -p 54321:5432 -e POSTGRES_PASSWORD=postgres grpc-db;

.PHONY: build_image
build_image:
	@docker build .  -t grpc-db;
