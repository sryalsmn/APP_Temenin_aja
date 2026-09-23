# Backend Dockerfile for TEMENIN
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download || true

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o temenin-server ./cmd/server/main.go

FROM alpine:3.19

WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/temenin-server .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./temenin-server"]
