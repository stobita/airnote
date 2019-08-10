migrate-create:
	cd deployments/development && docker-compose exec api goose -dir=./db/migrations create $(NAME) sql
migrate:
	cd deployments/development && docker-compose exec api go run cmd/airnote/airnote.go migrate
sqlboiler:
	cd deployments/development && docker-compose exec api sqlboiler mysql --wipe -c ./db/sqlboiler.toml -o ./internal/repository/rdb -p rdb

