FROM golang:1.26-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o myapp ./auth-service/cmd/ \
&& go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates
RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/auth-service/db/migrations /migrations
COPY --from=builder --chown=appuser:appuser /app/myapp /app/myapp
USER appuser
CMD goose -dir /migrations postgres "user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} port=${DB_PORT} host=${DB_HOST}" up && exec /app/myapp