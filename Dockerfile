# Start from the official Go image
FROM golang:1.23.2-alpine
# FROM golang:1.23.2-bullseye

# Set environment variable for Gin mode
ENV GIN_MODE=release
ENV APP_HOME /client

# Set the current Working Directory
WORKDIR ${APP_HOME}

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ${APP_HOME}/cmd/main ${APP_HOME}/cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/client/cmd/main"]