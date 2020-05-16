FROM golang:alpine as builder

LABEL author="wdalmijn <winbulk@gmail.com>"

COPY . /go/src/app

WORKDIR /go/src/app

CMD "go run main.go"