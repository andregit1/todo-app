# Use the official Golang image as the base image
FROM golang:1.22.4-alpine3.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Install necessary packages
RUN apk update && apk add --no-cache git

# Install swag for API documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000
