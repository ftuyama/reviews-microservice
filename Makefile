install:
	go mod init reviews
	go mod download
	go build -o app

build:
	go build -o app

run:
	./app -port=8080 -database=mongodb -mongo-host=localhost:27016

docker_up:
	docker-compose up -d reviews-db

docker_down:
	docker-compose down
