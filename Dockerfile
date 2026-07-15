# -------------------------
# Build Stage
# -------------------------

FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN go build -o main ./cmd/server

# -------------------------
# Production Stage
# -------------------------

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

EXPOSE 8080

CMD ["./entrypoint.sh"]