# Use the official Go image as a base image
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go files to the working directory
COPY . .

# Build the Go application
RUN go build -o labora-movies cmd/labora-movies

# Use a lightweight image as the final image
FROM alpine:latest

# Set the working directory in the new image
WORKDIR /app

# Copy the built binary from the builder image
COPY --from=builder /app/labora-movies .

# Expose the port on which your Go application runs
EXPOSE 8080

# Command to run your Go application
CMD ["./labora-movies"]
