# First stage: build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go files into the container
COPY . .

# Download dependencies and build the Go application
RUN go mod tidy
RUN go build -o travelapp

# Second stage: create a lightweight image
FROM alpine:latest

# Install necessary CA certificates for HTTPS communication
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/travelapp /travelapp

COPY --from=builder /app/configs /configs

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the executable
CMD ["./travelapp"]