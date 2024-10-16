FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
ARG APP_NAME
RUN go build -o /app/service.out ./cmd/${APP_NAME}/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service.out .
EXPOSE 8080
CMD ["./service.out"]