# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
ARG SERVICE_NAME
RUN go build -o main ./cmd/${SERVICE_NAME}

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./

# Expose port (will be overridden by docker-compose)
EXPOSE 8000

CMD ["./main"]