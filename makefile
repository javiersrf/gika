#WORKERS
run:
	go run main.go

build:
	go build -o ./build/consumer main.go

clean:
	docker system prune -f

db:
	docker-compose up --build -d