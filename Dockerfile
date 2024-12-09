FROM golang:1.22.3 as build
WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]