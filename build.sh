#!/bin/bash

# Exit script on any error
set -e

# Build and run Docker Compose services
docker-compose up -d

# Perform other actions:


# Clean up - stop and remove Docker Compose services
docker-compose down

#for now this will return 0 because nothing can be built as of yet.
exit 0