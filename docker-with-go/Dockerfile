FROM golang:1.12.0-alpine3.9

# Create app directory
RUN mkdir /app

# Copy application code to container
ADD . /app
WORKDIR /app

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
