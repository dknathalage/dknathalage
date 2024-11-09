FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/invoice

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main", "serve", "--http=0.0.0.0:80"]