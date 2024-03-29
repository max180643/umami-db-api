ARG PORT=8080
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
COPY --from=builder /app/IP2LOCATION-LITE-DB1.IPV6.BIN .

RUN apk update
RUN apk add --no-cache tzdata
RUN apk add --no-cache curl

HEALTHCHECK --interval=30s --timeout=15s --start-period=30s \
  CMD curl --silent --fail localhost:$PORT/health || exit 1

EXPOSE $PORT
CMD [ "/app/main" ]