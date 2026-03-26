.PHONY: up down test-client logs scale-2 scale-5

up:
	docker-compose up -d --build --scale payment-service=3

down:
	docker-compose down

test-client:
	go run client/main.go

logs:
	docker-compose logs -f

scale-2:
	docker-compose up -d --scale payment-service=2

scale-5:
	docker-compose up -d --scale payment-service=5
