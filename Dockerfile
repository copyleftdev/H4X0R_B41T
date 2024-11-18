# Use a lightweight base image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the Go source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o h4x0r_b41t

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./h4x0r_b41t"]
