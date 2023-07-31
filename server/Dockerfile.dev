FROM golang:1.19-buster as builder

# install dependencies
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

EXPOSE 8080
# Copy local code to the container image.
COPY . ./

CMD ["air", "-c", ".air.toml"]