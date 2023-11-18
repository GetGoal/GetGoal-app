FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container at /app
COPY . .

# ARG directive allows passing build-time variables, in this case, the ENV_FILE
ARG ENV

# Copy the environment-specific .env file
COPY .env.${ENV} .env

# Download and install any required third-party dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container to run the executable
CMD ["./main"]