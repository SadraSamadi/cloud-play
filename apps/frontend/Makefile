DOCKER_ORG ?= sadrasamadi
DOCKER_IMG ?= shuttle-frontend
DOCKER_TAG ?= latest

docker-publish:
	docker build -t $(DOCKER_ORG)/$(DOCKER_IMG):$(DOCKER_TAG) .
	docker push $(DOCKER_ORG)/$(DOCKER_IMG):$(DOCKER_TAG)
