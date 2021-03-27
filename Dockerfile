FROM golang:1.16-alpine

RUN mkdir -p /go/bin
RUN mkdir -p /go/pkg
RUN mkdir -p /go/src
RUN mkdir -p /app

RUN export GOPATH=$HOME/go
RUN export GOROOT=/usr/local/opt/go/libexec
RUN export PATH=$PATH:$GOPATH/bin
RUN export PATH=$PATH:$GOROOT/bin

WORKDIR /app

COPY ./ .

RUN go mod tidy

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]
