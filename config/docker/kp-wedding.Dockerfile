FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY pkg ./pkg
COPY templates ./templates
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o cmd/kp-wedding/main ./cmd/kp-wedding

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cmd/kp-wedding/main /app/cmd/kp-wedding/main
COPY --from=builder /app/templates ./templates
RUN ls -R /app/
EXPOSE 80
CMD ["./cmd/kp-wedding/main"]