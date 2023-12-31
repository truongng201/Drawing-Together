# ---------------------------------------
# Development stage
# ---------------------------------------
FROM debian:latest AS development

# Set environment variables for Go installation
ENV GOLANG_VERSION 1.18
ENV GOLANG_URL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz

# Install required dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates git curl && \
    rm -rf /var/lib/apt/lists/*

# Download and install Go
RUN wget -O go.tar.gz $GOLANG_URL && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GO111MODULE=on

# install air golang dependencies for devlement
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Create and change to the app directory.
WORKDIR /app

COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]