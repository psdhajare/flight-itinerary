# Dockerfile

# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/itinerary cmd/server/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bin/itinerary .

EXPOSE 8080

ENTRYPOINT ["./itinerary"]