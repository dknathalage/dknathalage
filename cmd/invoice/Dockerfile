# Step 1: Use the official Golang image as the builder
FROM golang:1.23-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the rest of the application code
COPY . .

# Step 5: Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/invoice

# Step 6: Use a minimal image to run the binary
FROM alpine:latest

# Step 7: Set the working directory
WORKDIR /app

# Step 8: Copy the binary from the builder image
COPY --from=builder /app/main .

# Step 9: Expose the application port
EXPOSE 8080

# Step 10: Run the application
CMD ["./main", "serve", "--http=0.0.0.0:80"]