FROM golang:1.21-alpine AS builder

#ARG ARG_HTTP_PROXY 127.0.0.1:1234
#ARG ARG_HTTPS_PROXY 127.0.0.1:1234
#RUN sed -i 's/https/http' /etc/apk/repositories

RUN apk update \
    && apk --no-cache --update add build-base gcc

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main.o /src/cmd/service/main.go

FROM alpine:3.18.2 AS production-stage
WORKDIR /src
RUN mkdir -p /src/logs
COPY --from=builder /src/main.o /src
CMD ["/src/main.o"]
