# develop command

.PHONY: run
gorun:
	air -c .air.toml

# docker compose command

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: build
build:
	docker compose build

.PHONY: ps
ps:
	docker compose ps

.PHONY: logs
logs:
	docker compose logs -f


