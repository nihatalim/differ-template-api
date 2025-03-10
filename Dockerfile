# Build stage
FROM golang:1.21.8-alpine AS builder

# Set working directory
WORKDIR /app

# Copy Go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application source code
COPY . .

# Build Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o differ-template-api ./cmd/api

# Run stage
FROM alpine:latest

# Add certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/differ-template-api .

# Copy configs and other necessary files
COPY --from=builder /app/api/ /app/api/

# Run the application
CMD ["./differ-template-api"] 