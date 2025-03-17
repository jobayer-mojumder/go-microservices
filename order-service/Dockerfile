# Use Golang base image
FROM golang:1.24

# Set working directory
WORKDIR /app

# get dependencies for live reloading (air)
RUN go install github.com/air-verse/air@latest

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Expose the application port
EXPOSE 8084

# Use air to run the application in development mode (with live-reloading)
CMD ["air"]
