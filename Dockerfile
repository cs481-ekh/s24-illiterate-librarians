# Use a multi-stage build for smaller and more secure images
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to cache dependencies
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY backend/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o my-golang-app .

# Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/my-golang-app .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the compiled binary
CMD ["./my-golang-app"]

# Use a multi-stage build for smaller and more secure images
FROM node:latest AS frontend-builder

WORKDIR /app

COPY frontend/package*.json ./
RUN npm install

COPY frontend/ .
RUN npm run build

