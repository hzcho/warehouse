FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app/
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/app/main .

ENTRYPOINT ["./main"]