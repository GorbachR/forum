startDb:
	cd config/docker && docker compose up -d

stopDb:
	cd config/docker && docker compose down

migrateUp:
	@echo "How many migrations do you want to apply?"
	@read num_migrations; \
	migrate -source file://db/migrations -database postgres://postgres:dev@localhost:5432/forum?sslmode=disable up $$num_migrations

migrateDown:
	@echo "How many migrations do you want to apply?"
	@read num_migrations; \
	migrate -source file://db/migrations -database postgres://postgres:dev@localhost:5432/forum?sslmode=disable down $$num_migrations

sqlcgen:
	cd db && sqlc generate