OBJECTS=server.out

IMG=kafkatool-be
IMG_TAG=1.0.0

CONTAINER_REGISTRY = docker.io
USER = duyle95

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(OBJECTS) cmd/service/main.go

.PHONY: run
run:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(OBJECTS) cmd/service/main.go && ./$(OBJECTS)

.PHONY: docker-build
docker-build:
	docker build -t $(IMG):$(IMG_TAG) .

.PHONY: docker-push
docker-push:
	docker image tag $(IMG):$(IMG_TAG) $(CONTAINER_REGISTRY)/$(USER)/$(IMG):$(IMG_TAG)
	docker push $(CONTAINER_REGISTRY)/$(USER)/$(IMG):$(IMG_TAG)

.PHONY: clean
clean:
	rm $(OBJECTS)
