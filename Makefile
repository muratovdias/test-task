run:
	docker compose up
down:
	docker compose down
test-app1:
	go test ./... -v