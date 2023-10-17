OBJECTS=server.out

IMG=kafkatool
IMG_TAG=v1

CONTAINER_REGISTRY = docker.io
USER = duyle95

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: build
build:
	go build -o $(OBJECTS) cmd/service/main.go

.PHONY: docker-build
docker-build:
	docker build -t $(IMG):$(IMG_TAG) .

.PHONY: docker-push
docker-push:
	docker push $(CONTAINER_REGISTRY)/$(USER)/$(IMG):$(IMG_TAG)

.PHONY: clean
clean:
	rm $(OBJECTS)
