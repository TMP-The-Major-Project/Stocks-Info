.PHONY: all clean test build up down restart logs ps

all: restart

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

restart: down up

logs:
	docker-compose logs -f

ps:
	docker ps

