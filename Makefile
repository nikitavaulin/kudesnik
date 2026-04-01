include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	docker compose up -d kudesnik-postgres

env-down:
	docker compose down kudesnik-postgres
	
env-cleanup:
	@read -p "ATTENTION! Clear all volume env files? Lose data danger. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down kudesnik-postgres && \
		rm -rf out/pgdata && \
		echo "Env files were cleared"; \
	else \
		echo "Clearing was cancelled"; \
	fi