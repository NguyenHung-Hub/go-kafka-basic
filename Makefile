up:
	docker-compose up --build -d
down:
	docker-compose down
logs:
	docker-compose logs -f
top:
	docker stats
con:
	go run cmd/consumer/main.go
pro:
	go run cmd/producer/main.go


.PHONY: up down logs top con pro