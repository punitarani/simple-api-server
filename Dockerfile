# Build and Run Server with Docker

# Use Go 1.19
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Verify dependencies have expected content
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run the executable
CMD ["./main"]

# Expose port 8080 to the outside world
EXPOSE 8080

# Build the Docker image
# docker build -t simple-api-server .

# Run the Docker image with 4 workers
# docker run -p 8080:8080 --name simple-api-server simple-api-server -w 4

# Run the Docker image in the background
# docker run -d -p 8080:8080 --name simple-api-server simple-api-server -w 4
