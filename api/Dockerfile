FROM golang:1.15

ENV TZ='Asia/Tokyo'
ENV LANG=ja_JP.UTF-8

WORKDIR /api

COPY ./ /api

RUN apt-get update && apt-get install bash \
    && go mod download -x \
    && go get github.com/githubnemo/CompileDaemon \
    && chmod 755 ./main.go

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="go build -v main.go" --command=./main
