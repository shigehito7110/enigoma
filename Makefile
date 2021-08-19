up:
	docker-compose up -d
build:
	docker-compose build --no-cache --force-rm
down:
	docker-compose down --remove-orphans
ps:
	docker-compose ps
restart:
	@make down
	@make build
	@make up
start:
	@make build
	@make up
app:
	docker-compose exec app bash
log:
	docker-compose logs -f
