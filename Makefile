migrate-create:
	cd deployments/development && docker-compose exec api goose -dir=./db/migrations create $(NAME) sql
migrate:
	cd deployments/development && docker-compose exec api go run cmd/airnote/airnote.go migrate
