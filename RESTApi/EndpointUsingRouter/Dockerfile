# Light weight Go alpine image to work with
FROM golang:1.11.1 AS builder

# This is required for the mux router
RUN go get -u github.com/gorilla/mux
# create an /app directory in which
# we'll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app
# We want to build our application's binary executable
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

# The lightweight scratch image we'll run our application within
FROM alpine:latest AS production
# Copy the output from our builder stage to our production stage
COPY --from=builder /app .
# start the app using the newly compiled binary
CMD ["./main"]


