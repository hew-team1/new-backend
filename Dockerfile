# Dockerfileはビルドされる時に実行される内容

FROM golang:1.17
RUN go get github.com/cespare/reflex
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

