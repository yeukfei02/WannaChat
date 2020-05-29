FROM golang:1.14.2

RUN mkdir -p /go/bin
RUN mkdir -p /go/pkg
RUN mkdir -p /go/src
RUN mkdir -p /go/src/WannaChat

RUN export GOPATH=$HOME/go
RUN export GOROOT=/usr/local/opt/go/libexec
RUN export PATH=$PATH:$GOPATH/bin
RUN export PATH=$PATH:$GOROOT/bin

WORKDIR /go/src/WannaChat

COPY ./ .

RUN go mod tidy

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]
