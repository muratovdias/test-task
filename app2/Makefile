IMAGE_NAME=app2
CONTAINER_NAME=app2-container
PORT=8888

.PHONY: build run

build:
	sudo docker build -t $(IMAGE_NAME) .
run-mongo:
	sudo docker run --name my-mongodb --add-host=my-mongodb:172.19.0.1 --network my-network -p 27017:27017 mongo
run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) --network my-network -p 8888:8888 $(IMAGE_NAME)

