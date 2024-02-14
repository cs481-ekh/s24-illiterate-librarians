# Use an official Node.js image as the base image
FROM node:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files to the container
COPY . .

# Build the Angular app
RUN npm run build

# Use an official Apache image as the base image for serving the app
FROM httpd:latest

WORKDIR /usr/local/apache2/htdocs/

# Copy the built Angular app from the build stage to the Apache document root
COPY --from=build /app/dist/frontend/browser .

# Expose port 80 for Apache
EXPOSE 80

# Start Apache in the foreground
CMD ["httpd-foreground"]