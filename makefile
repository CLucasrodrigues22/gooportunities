.PHONY: default run build test docs clean
# Variables
APP_NAME=gopportunities

# Tasks
default: run-docker-up

run-docker-up:
	@swag init
	@docker-compose up -d
	@docker ps
run-docker-down:
	@docker-compose down
run:
	@go run main.go
run-local:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs