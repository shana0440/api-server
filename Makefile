MIGRATIONS=db/migrations
DSN=postgresql://\$$DB_USER:\$$DB_PASSWORD@\$$DB_HOST:\$$DB_PORT/\$$DB_NAME?sslmode=disable

create-migration:
	docker-compose run --rm app bash -c "migrate -verbose create -ext sql -dir $(MIGRATIONS) $(NAME)"

migrate-up:
	docker-compose run --rm app bash -c "migrate -verbose -path $(MIGRATIONS) -database $(DSN) up"

migrate-down:
	docker-compose run --rm app bash -c "migrate -verbose -path $(MIGRATIONS) -database $(DSN) down 1"