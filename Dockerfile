FROM gcr.io/cloud-builders/docker

FROM golang:latest

LABEL maintainer "Brodie <brodiep21@hotmail.com>"

WORKDIR /NGIC

COPY index.html .

COPY main.go .

COPY go.mod .

COPY go.sum .

RUN mkdir assets

COPY assets/calc.js assets/

COPY assets/styles.css assets/

ENV PORT 8080

# RUN apk add --no-cache go

# RUN go version

# RUN go get github.com/rocketlaunchr/google-search

RUN go build

ENTRYPOINT go run main.go