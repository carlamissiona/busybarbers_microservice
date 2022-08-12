FROM golang:1.18-alpine

ENV GO111MODULE=on

RUN mkdir -p /go/src/

ENV APP_HOME /go/src/
 
WORKDIR /go/src

COPY . .

RUN go mod tidy

RUN go build -o /go/src/main

RUN export PATH="/go/src/main:$PATH"

EXPOSE 8000

ENTRYPOINT ["/go/src/main","api"]
 
