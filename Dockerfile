# # Use an official Node.js image as the base image
FROM node:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container from the front end folder
COPY /frontend/package*.json .

# Install dependencies
RUN npm install


RUN npm install -g @angular/cli

# Copy the current directory contents into the container at /app
COPY /frontend .

# Build the Angular app
RUN ng build --configuration=production

# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory in the container
ENV ROUTE_PREFIX=/s24-illiterate-librarians

# Set the working directory to /backend
WORKDIR /backend

# Copy the current directory contents into the container at /backend
COPY /backend /backend

COPY --from=build /app/dist/frontend/browser .

RUN go get
# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]