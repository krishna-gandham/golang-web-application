# Use official Golang image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the Go code into the container
COPY main.go .

# Build the Go application
RUN go build -o main main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
