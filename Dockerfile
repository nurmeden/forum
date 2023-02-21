FROM golang:latest

WORKDIR /forum

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]