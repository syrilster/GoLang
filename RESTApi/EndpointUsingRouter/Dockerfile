FROM golang:1.11.1 AS builder

RUN go get -u github.com/gorilla/mux
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]


