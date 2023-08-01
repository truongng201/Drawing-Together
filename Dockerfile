# Use Debian as the base image
FROM debian:latest

# Set environment variables for Go installation
ENV GOLANG_VERSION 1.18
ENV GOLANG_URL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz

# Set environment variables for Node.js installation
ENV NODE_VERSION 18.x

# Install required dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates git && \
    rm -rf /var/lib/apt/lists/*

# Download and install Go
RUN wget -O go.tar.gz $GOLANG_URL && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# Download and install Nodejs
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash - && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GO111MODULE=on

# Verify Go installation
RUN go version

# Verify Node.js and npm versions
RUN node -v && npm -v

# Set the working directory
WORKDIR /app

# Add non-root user
RUN addgroup --system user && adduser --system --no-create-home --group user
RUN chown -R user:user /app && chmod -R 755 /app

# Switch to non-root user
USER user

COPY . .

# Install dependencies for the server
RUN cd server && go mod download

# Install dependencies for the client
RUN cd client && npm install

# Build the client
RUN cd client && npm run build

# Build the Go app
RUN cd server/cmd && go build -o server

# Command to run the executable
CMD cd server/cmd && ./server
