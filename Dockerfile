# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/jtihinjk/plivoapp

# Build the command inside the container.
RUN go install github.com/jithinjk/plivoapp

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/plivoapp

# http server listens on port 8080.
EXPOSE 8080
