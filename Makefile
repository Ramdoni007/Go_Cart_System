postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root cart_system

dropdb:
	docker exec -it postgres15 dropdb cart_system

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cart_system?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cart_system?sslmode=disable" -verbose down

migratefix:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/cart_system?sslmode=disable" force 1
.PHONY: postgres createdb dropdb