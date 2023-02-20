FROM golang:alpine AS builder

WORKDIR /forum/

COPY . .

RUN go build -o main ./cmd/

FROM alpine

WORKDIR /forum/

LABEL "author"="dulat"
LABEL build_date="2023-02-21"

COPY --from=builder /forum/ /forum/

EXPOSE 8080

CMD ["./cmd"]