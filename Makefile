SERVICE_NAME := go-backend-template
DB_NAME := go-backend-template
DB_CONNECTION := postgresql://postgres:postgres@localhost:5432/$(DB_NAME)?sslmode=disable

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres $(DB_NAME)

dropdb:
	docker exec -it postgres dropdb --username=postgres $(DB_NAME)

install:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
	go install github.com/google/wire/cmd/wire@latest && \
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest && \
	go install github.com/golang/mock/mockgen@latest && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
	export PATH="$PATH:$(go env GOPATH)/bin"

compile: api sqlc proto wire

api:
	oapi-codegen -package api -generate gin,types,spec -exclude-tags="actuator" swagger/api.yaml > api/oapi.go

wire:
	wire ./...

migrateup:
	migrate -path db/migration -database $(DB_CONNECTION) -verbose up

migrateup1:
	migrate -path db/migration -database $(DB_CONNECTION) -verbose up 1

migratedown:
	migrate -path db/migration -database $(DB_CONNECTION) -verbose down

migratedown1:
	migrate -path db/migration -database $(DB_CONNECTION) -verbose down 1

sqlc:
	rm -rf repository/sqlc && sqlc generate

proto:
	rm -rf proto/pb
	mkdir -p proto/pb
	protoc --proto_path=proto proto/*.proto --go_out=proto --go-grpc_out=proto

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go $(SERVICE_NAME)/db/sqlc Store

.PHONY: createdb dropdb install compile api migrateup migrateup1 migratedown migratedown1 sqlc proto test server mock wire
