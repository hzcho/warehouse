FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app/
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY ./private_key.pem ./private_key.pem
COPY ./public_key.pem ./public_key.pem
COPY --from=builder /app/cmd/app/main .

ENTRYPOINT ["./main"]