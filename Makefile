proto:
	protoc --proto_path=./pkg --go-grpc_out=. --go_out=paths=import:. ./pkg/pb/*.proto

server:
	go run cmd/main.go

db:
	cd internal/db && \
		docker-compose up -d
