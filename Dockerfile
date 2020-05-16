FROM golang:alpine as builder

LABEL author="wdalmijn <winbulk@gmail.com>"

COPY . /go/src/app

WORKDIR /go/src/app

RUN go build

CMD ./golang-xml-parse-bench