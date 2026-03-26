.PHONY: up down test-client logs

up:
	docker-compose up -d --build --scale payment-service=3

down:
	docker-compose down

test-client:
	go run client/main.go

logs:
	docker-compose logs -f

proto-install:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/service.proto