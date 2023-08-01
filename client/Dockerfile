# Use Debian as the base image
FROM debian:latest

# Set environment variables for Node.js installation
ENV NODE_VERSION 18.x

# Install required dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Download and install Nodejs
RUN wget -qO- https://deb.nodesource.com/setup_${NODE_VERSION} | bash - && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Verify Node.js and npm versions
RUN node -v && npm -v

# Set the working directory
WORKDIR /app

# Copy application to working directory
COPY . .

# Install dependencies for the client
RUN cd client && npm install

# Build the client
RUN cd client && npm run build

# Command to run the executable
CMD cd client && npm run start

