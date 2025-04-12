# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notion-htmx-blog ./cmd/server

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/notion-htmx-blog .

# Copy static files and templates
COPY --from=builder /app/web/static ./web/static
COPY --from=builder /app/web/templates ./web/templates

# Expose port
EXPOSE 8080

# Run the application
CMD ["./notion-htmx-blog"] 