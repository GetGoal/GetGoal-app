FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container at /app
COPY . .

# ARG directive allows passing build-time variables, in this case, the ENV_FILE
ARG ENV
ARG DB_HOST
ARG DB_USER
ARG DB_PORT
ARG DB_NAME
ARG DB_PASSWORD

# Set environment

ENV ENV=${ENV}
ENV DATABASE_HOST=${DB_HOST}
ENV DATABASE_PORT=${DB_PORT}
ENV DATABASE_USER=${DB_USER}
ENV DATABASE_PASSWORD=${DB_PASSWORD}
ENV DATABASE_DBNAME=${DB_NAME}
ENV DATABASE_SSLMODE=disable
ENV DATABASE_TIMEZONE=Asia/Bangkok

# Download and install any required third-party dependencies
RUN go mod download

RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container to run the executable
CMD ["./main"]