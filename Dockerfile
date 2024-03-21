# two stage build

FROM docker.io/golang:1.22.0 AS build

WORKDIR /app

# Download swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy app to container
COPY app .

# Generate swagger docs
RUN swag init --pd --parseDepth 3

# cgo is needed for sqlite
RUN go mod download && go test ./... && CGO_ENABLED=1 GOOS=linux go build -o backend

# we need a base image for dynamic linked libs and can't work from scratch
FROM docker.io/debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y ca-certificates && \ 
    rm -rf /var/lib/{apt,dpkg,cache,log}/

ENV ENV dev
ENV API_DB_CONNECTION_STR "sqlite:///app/data/data.db"

EXPOSE 9000
VOLUME ["/app/data"]
VOLUME ["/app/config"]

WORKDIR /app

COPY --from=build /app/backend backend

ENTRYPOINT ["/app/backend"]
