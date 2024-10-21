FROM golang:1.23-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /run-app ./cmd/service


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]