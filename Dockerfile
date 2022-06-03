FROM gcr.io/cloud-builders/docker

FROM golang:latest

LABEL maintainer "Brodie <brodiep21@hotmail.com>"

WORKDIR /NGIC

COPY go.mod .

COPY go.sum .

COPY assets/calc.js .

COPY assets/styles.css .

COPY index.html .

COPY main.go .

ENV PORT 8080

# RUN apk add --no-cache go

# RUN go version

RUN go build

ENTRYPOINT go run main.go