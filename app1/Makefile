IMAGE_NAME=app1
CONTAINER_NAME=app1-container
PORT=8887

.PHONY: build run

build:
	sudo docker build -t $(IMAGE_NAME) .

run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) --network my-network $(IMAGE_NAME)

