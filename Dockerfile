# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container at /app
COPY . .

# Download and install any required third-party dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container to run the executable
CMD ["./main"]