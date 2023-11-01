# two stage build

FROM docker.io/golang:1.21.3 AS build

WORKDIR /app

COPY app .

# maybe cgo is needed for sqlite
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o backend

FROM scratch

WORKDIR /app

COPY --from=build /app/backend backend

ENTRYPOINT ["/app/backend"]