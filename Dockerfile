FROM golang:buster

# Environment variables.
ENV GOPATH=$GOPATH:/usr/share/gocode

# Install packages from APT.
RUN apt-get update
RUN apt-get upgrade -y
RUN go get github.com/go-telegram-bot-api/telegram-bot-api

WORKDIR /usr/local/go/src/github.com/aniolm9/ninini-bot
COPY . /usr/local/go/src/github.com/aniolm9/ninini-bot

RUN go get ./
RUN go build

CMD go get github.com/pilu/fresh && fresh
