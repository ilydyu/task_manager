run:
	go run ./cmd/main.go

up:
	docker compose up

down:
	docker compose down

.PHONY: test
test:
	go test -v -cover -coverpkg=./... ./test/integration