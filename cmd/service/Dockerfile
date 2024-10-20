FROM golang:1.23-bookworm as builder

ARG PATH

WORKDIR /usr/src/app
COPY ../../go.mod ../../go.sum ./
RUN go mod download && go mod verify
COPY ../../ .
RUN go build -v -o /run-app $PATH


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]