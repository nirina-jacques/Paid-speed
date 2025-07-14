cd.PHONY: help
.DEFAULT_GOAL = help

##@ Docker
build_p: ## Construire les containers docker
	docker-compose build

start_p: ## Lancer les containers docker
	echo docker-compose up -d --remove-orphans
	docker-compose up
exec_p: ## Lancer les containers docker
	echo docker-compose up --remove-orphans
	docker exec -it app_front bash
stop_p:
	docker-compose down