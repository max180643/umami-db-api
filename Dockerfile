# Build stage
FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -o main src/app.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

RUN apk update
RUN apk add curl

HEALTHCHECK --interval=30s --timeout=15s --start-period=30s \
  CMD curl --silent --fail localhost:8080/health || exit 1

EXPOSE 8080
CMD [ "/app/main" ]