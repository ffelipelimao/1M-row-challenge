# Stage 1: Build the Go application
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

# Stage 2: Create a minimal Docker image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api ./api

# Expose the port the application will run on
EXPOSE 8080

# Make the binary executable
RUN chmod +x ./api

# Command to run the application
CMD ["./api"]