. phony: build
build: 
	go build -v ./cmd/api
.DEFAULT_GOAL := build
POST_DRIVERS := "postgres"
POSTGRES_DBSTRING := "host=localhost port=5432 user=aleksandr password=pwd1234 dbname=hw14_sql sslmode=disable"

connect_db:
	@docker exec -it hw14_sql psql -U aleksandr

show-user-table:
	@docker exec -it hw14_sql psql -U aleksandr -c "SELECT NOW(); SELECT * FROM user;"


migrate-up:
	@goose -dir internal/migrations $(POST_DRIVERS) $(POSTGRES_DBSTRING) up

migrate-down:
	@goose -dir internal/migrations $(POST_DRIVERS) $(POSTGRES_DBSTRING) down

create-db:
	createdb -h localhost -p 5432 -U $(POSTGRES_USER) $(POSTGRES_DB)