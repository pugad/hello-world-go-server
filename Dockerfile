FROM golang:1.18.0-alpine

WORKDIR /usr/src/helloworld
COPY helloworld.go .
ENTRYPOINT ["go","run","helloworld.go"]