
FROM golang:1.12.9-alpine3.9 AS builder
LABEL maintainer=adedayoakinpelu@gmail.com
RUN go version

COPY . /go/src/github.com/haroldsphinx/go-coding-challenge/
WORKDIR /go/src/github.com/haroldsphinx/go-coding-challenge/


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .


EXPOSE 8083
ENTRYPOINT ["./app"]