build:
	go build -o bin/app cmd/api/main.go

run: 
	go run cmd/api/main.go

test:
	go test ./... -cover

docker-build:
	docker build -t multiplayer-modes-app . 

docker-run: 
	docker run -p 8080:8080 --env-file .env multiplayer-modes-app

docker-compose-up: 
	docker-compose up --build 

docker-compose-down:
	docker-compose down