FROM golang:1.22.3-alpine

WORKDIR /app

ENV GO111MODULE=on

COPY . . 

RUN go mod download

RUN go build -o main ./cmd

EXPOSE 8081

CMD ["./main"]
