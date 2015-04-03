# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

MAINTAINER George Sun <http://codethoughts.info>

# For convenience, set an env variable with the path of the code
ENV APP_DIR  $GOPATH/src/github.com/puffsun/todos

# Copy the local package files to the container's workspace.
ADD . $APP_DIR

WORKDIR $GOPATH/src/github.com/puffsun/todos

# Install dependencies
RUN go get github.com/tools/godep
RUN go get github.com/pilu/fresh
RUN $GOPATH/bin/godep restore

# Build the project inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/puffsun/todos

# Run the web application by default when the container starts.
ENTRYPOINT $GOPATH/bin/todos

# Document that the service listens on port 8000.
EXPOSE 8000
