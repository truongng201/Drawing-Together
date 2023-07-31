# Use Debian as the base image
FROM debian:latest

# Set environment variables for Go installation
ENV GOLANG_VERSION 1.17
ENV GOLANG_URL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz

# Install required dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Download and install Go
RUN wget -O go.tar.gz $GOLANG_URL && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GO111MODULE=on

# Verify Go installation
RUN go version

# Set the working directory
# Create a directory for the Go workspace
RUN mkdir -p $GOPATH/src $GOPATH/bin

# Set the working directory
WORKDIR $GOPATH/src

COPY . .

RUN cd server && go mod download
# Build the Go app
RUN cd server && go build cmd/server.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["cd", "server", "&&", "cmd/server"]
