NAME=atlant
DEV_DOCKER_COMPOSE=dockerfiles/dev/docker-compose.yml

.PHONY: dev_env_up dev_env_down dev_env_migrate dev_env_deploy
dev_env_up: COMPOSE_FILE=$(DEV_DOCKER_COMPOSE)
dev_env_up: env_up
dev_env_down: COMPOSE_FILE=$(DEV_DOCKER_COMPOSE)
dev_env_down: env_down
dev_env_migrate: COMPOSE_FILE=$(DEV_DOCKER_COMPOSE)
dev_env_migrate: env_migrate
dev_env_deploy: COMPOSE_FILE=$(DEV_DOCKER_COMPOSE)
dev_env_deploy: env_deploy
dev_env_deploy_broker: COMPOSE_FILE=$(DEV_DOCKER_COMPOSE)

env_up: env_down
	time docker-compose -f $(COMPOSE_FILE) up --build $(NAME)

env_down:
	time docker-compose -f $(COMPOSE_FILE) down

env_migrate:
	time docker-compose -f $(COMPOSE_FILE) up --build migrations

env_deploy:
	time docker-compose -f $(COMPOSE_FILE) up -d --no-recreate --build $(NAME)