
.PHONY: db/init
db/init:
	docker compose exec db mysql -u guild-hack -pguild-hack -e'show databases;'

.PHONY: flyway/info
flyway/info:
	docker compose run --rm flyway info

