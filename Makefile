include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d kudesnik-postgres

env-down:
	@docker compose down kudesnik-postgres
	
env-cleanup:
	@read -p "ATTENTION! Clear all volume env files? Lose data danger. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down kudesnik-postgres && \
		rm -rf out/pgdata && \
		echo "Env files were cleared"; \
	else \
		echo "Clearing was cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Argument seq is missed. Want: make migrate-create seq=init";\
		exit 1; \
	fi; \
	docker compose run --rm kudesnik-postgres-migrate \
		create \
		-ext sql \
		-dir //migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Argument action is missed. Want: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm kudesnik-postgres-migrate \
		-path //migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@kudesnik-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		"$(action)"

kudesnik-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run cmd/kudesnik/main.go