FROM golang:latest

WORKDIR /app
ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download
ADD . /app
RUN go build -o app

ENTRYPOINT [ "./app" ]