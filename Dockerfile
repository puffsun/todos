# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/puffsun/todos

# Install dependencies TODO
RUN go get

# Build the project inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/puffsun/todos

# Run the web application by default when the container starts.
ENTRYPOINT /go/bin/todos

# Document that the service listens on port 8000.
EXPOSE 8000
