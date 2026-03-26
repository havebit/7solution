run:
	go run main.go

init-db:
	bash setup_db.sh

build:
	go build -o app main.go

docker-build:
	docker build -t todo-api .

docker-run:
	docker run -p 8080:8080 todo-api

clean:
	rm -f todo.db app
