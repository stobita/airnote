DEV_COMPOSE=deployments/development/docker-compose.yml
TEST_COMPOSE=deployments/test/docker-compose.yml

dev-up:
	docker-compose -f $(DEV_COMPOSE) up -d
dev-down:
	docker-compose -f $(DEV_COMPOSE) down
dev-logs:
	docker-compose -f $(DEV_COMPOSE) logs -f
migrate-create:
	docker-compose -f $(DEV_COMPOSE) exec api goose -dir=./db/migrations create $(NAME) sql

migrate:
	docker-compose -f $(DEV_COMPOSE) exec api go run cmd/airnote/airnote.go migrate
sqlboiler:
	docker-compose -f $(DEV_COMPOSE) exec api sqlboiler mysql --wipe -o ./internal/repository/rdb -c ./db/sqlboiler.toml -p rdb --no-auto-timestamps --no-tests

test-go:
		docker-compose -f $(TEST_COMPOSE) up --abort-on-container-exit && \
		docker-compose -f $(TEST_COMPOSE) down --volumes
