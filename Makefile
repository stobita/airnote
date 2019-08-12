migrate-create:
	cd deployments/development && docker-compose exec api goose -dir=./db/migrations create $(NAME) sql
migrate:
	cd deployments/development && docker-compose exec api go run cmd/airnote/airnote.go migrate
sqlboiler:
	cd deployments/development && docker-compose exec api sqlboiler mysql --wipe -o ./internal/repository/rdb -c ./db/sqlboiler.toml -p rdb --no-auto-timestamps
go-test:
	cd deployments/test && \
		docker-compose up --abort-on-container-exit && \
		docker-compose down --volumes
