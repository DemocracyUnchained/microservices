.PHONY: open init_db reset start stop validate

open: ## Open a shell into the web container
	docker exec -it democracy_web_1 env TERM=$(TERM) /bin/bash

init_db: ## Initialize the database.
	./init_db.sh

delete_db: ## Delete all database files. NOT KIDDING!
	sudo rm -rf .dev/mysql/*

reset: ## Remove all containers
	docker-compose down

start: ## Start all containers
	docker-compose up

stop: ## Stop all containers
	docker-compose stop

validate: ## Validate the compose configuration
	docker-compose config
