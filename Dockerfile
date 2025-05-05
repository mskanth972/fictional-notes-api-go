# Start from official Go image
FROM golang:1.21-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the application port
EXPOSE 8000

# Command to run the binary
CMD ["./main"]
