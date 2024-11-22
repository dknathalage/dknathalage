FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd ./cmd
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./modules/invoice/cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"]