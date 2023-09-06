FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /app/main.go


# Multi-stage build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

EXPOSE 23

CMD [ "/app/main" ]
