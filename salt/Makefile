IMAGE_NAME=app1
CONTAINER_NAME=app1-container
PORT=8887

.PHONY: build run

build:
	sudo docker build -t $(IMAGE_NAME) .

run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) --network my-network -p 8887:8887 $(IMAGE_NAME)

test:
	go test ./internal/service -v
	go test ./internal/handler -v