APP_NAME ?= shuttle
IMAGE_NAME ?= backend
IMAGE_TAG ?= latest
SERVER_PORT ?= 8080

docker-build:
	docker build -t $(APP_NAME)/$(IMAGE_NAME):$(IMAGE_TAG) .

docker-run:
	docker run --rm -p $(SERVER_PORT):8080 $(APP_NAME)/$(IMAGE_NAME):$(IMAGE_TAG)
