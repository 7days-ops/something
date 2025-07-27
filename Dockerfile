FROM golang:1.24.2 AS builder

WORKDIR /app

COPY . . 

RUN go build -o myapp:latest . 

FROM alpine:latest

WORKDIR /myapp

COPY --from=builder /myapp/app .

CMD [".myapp"]

