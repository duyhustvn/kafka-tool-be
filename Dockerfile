FROM golang:1.24 AS builder

ARG ARG_HTTP_PROXY
ARG ARG_HTTPS_PROXY

# RUN sed -i 's/https/http' /etc/apk/repositories

RUN apt update \
    && apt install -y gcc sqlite3 libsqlite3-dev

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main /src/cmd/service/main.go

FROM debian:bookworm-slim AS production-stage
WORKDIR /src

ARG ARG_HTTP_PROXY
ARG ARG_HTTPS_PROXY

RUN mkdir -p /src/logs
RUN mkdir -p /src/data

COPY --from=builder /src/main ./kafkatool-be

CMD ["./kafkatool-be"]
