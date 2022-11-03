IMAGE_ORG ?= tylergu1998
IMAGE_PROJECT ?= zkapp
IMAGE_TAG ?= v1
IMAGE_NAME ?= $(IMAGE_ORG)/$(IMAGE_PROJECT):$(IMAGE_TAG)

docker-image:
	docker build -t $(IMAGE_NAME) .

docker-push:
	docker push $(IMAGE_NAME)