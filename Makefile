# name app
# APP_NAME = server

# run:
#   go run ./cmd/${APP_NAME}/

# Tên của ứng dụng của bạn
APP_NAME := server

# Chạy ứng dụng
dev:
	go run ./cmd/$(APP_NAME)
run:
	docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

.PHONY: run
