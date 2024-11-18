FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY pkg ./pkg
COPY templates ./templates
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/invoice

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
EXPOSE 80
CMD ["./main"]