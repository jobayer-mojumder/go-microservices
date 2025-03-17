# Use Golang base image
FROM golang:1.24.1-alpine

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application
RUN go build -o user-service

# Expose the application port
EXPOSE 8081

# Run the application
CMD ["/app/user-service"]
