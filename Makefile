SHELL=/bin/bash
export REPO=grpc

.PHONY: run-local
run_local:
	@docker run -p 54321:5432 grpc-db -e POSTGRES_PASSWORD=postgres;



.PHONY: build_image
build_image:
	@docker build .  -t grpc-db;
