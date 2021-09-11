# Dockerfileはビルドされる時に実行される内容

FROM golang:alpine as build-reflex
RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    rm -rf /var/cache/apk/*

RUN go get github.com/cespare/reflex

# dev, builder
FROM golang:1.17 AS golang
WORKDIR /app

# dev
FROM golang as dev
COPY --from=build-reflex /go/bin/reflex /go/bin/reflex

# builder
FROM golang AS builder
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /build/guild-hack-api-linux .

# release
FROM alpine AS app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /app/build/guild-hack-api /usr/local/bin/guild-hack-api
EXPOSE 8080

# ↓で実行しているファイルはlinuxようではないけどいいの？
# ↓もしこれでいいのであれば下記の形式のbuildを行わないと現状ない
ENTRYPOINT ["guild-hack-api"]
