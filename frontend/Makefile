APP_NAME ?= shuttle
IMAGE_NAME ?= frontend
IMAGE_TAG ?= latest
WWW_ROOT ?= ./www
SERVER_PORT ?= 3000

docker-build:
	docker build -t $(APP_NAME)/$(IMAGE_NAME):$(IMAGE_TAG) .

docker-run:
	docker run --rm \
 		-v $(PWD)/$(WWW_ROOT):/usr/share/nginx/html:ro \
 		-p $(SERVER_PORT):80 \
 		$(APP_NAME)/$(IMAGE_NAME):$(IMAGE_TAG)
