.PHONY: open init_db reset start stop validate

open: ## Open a shell into the web container
	docker exec -it docker_web_1 env TERM=$(TERM) /bin/bash

init_db: ## Initialize the database.
	./init_db.sh

reset: ## Remove all containers
	docker-compose down

start: ## Start all containers
	docker-compose up

stop: ## Stop all containers
	docker-compose stop

validate: ## Validate the compose configuration
	docker-compose config
