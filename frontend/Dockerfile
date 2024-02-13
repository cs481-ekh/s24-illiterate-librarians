# Use an official Node runtime as a parent image
FROM node:latest

# Set the working directory to /frontend
WORKDIR /frontend

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Install app dependencies
RUN npm install

# Copy the current directory contents into the container at /frontend
COPY . /frontend

# Expose port 4200 to the outside world
EXPOSE 4200

# Command to run the Angular app
CMD ["ng", "build"]
