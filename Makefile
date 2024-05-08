install:
	go mod init reviews
	go mod download
	go build -o app

build:
	go build -o app

run:
	./app -database=mongodb -mongo-host=localhost:27016

docker_db:
	docker-compose up -d reviews-db

drun:
	docker-compose up

docker_down:
	docker-compose down
