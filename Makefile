create_db:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

postgres_start:
	docker run --name postgres15 -p 5437:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
	
postgres_stop:
	docker stop postgres15
	docker rm postgres15

migrate_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...