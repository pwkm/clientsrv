# Start from the official Go image
FROM golang:latest

# FROM golang:1.23.2-bullseye

# Set the current Working Directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.* ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o server ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./server"]