migrate create -ext sql -dir misc/migrations CreateAuthor

migrate -path misc/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" up
migrate -path misc/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" down
