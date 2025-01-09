Migrate scripts for database migration

testing golang for prod app

migrate create -ext sql -dir misc/migrations -seq <NAME>

migrate -path misc/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" up
migrate -path misc/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" down
