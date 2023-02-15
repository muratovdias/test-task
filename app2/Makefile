IMAGE_NAME=app2
CONTAINER_NAME=app2-container
PORT=8888

.PHONY: build run

build:
	sudo docker build -t $(IMAGE_NAME) .
run-mongo:
	sudo docker run --name my-mongodb --network my-network mongo
run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) --network my-network $(IMAGE_NAME)

