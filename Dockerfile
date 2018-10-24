FROM golang:1.11.1

RUN mkdir -p /go/bin
RUN mkdir -p /go/pkg
RUN mkdir -p /go/src
RUN mkdir -p /go/src/WannaChat

RUN export GOPATH=$HOME/go
RUN export GOROOT=/usr/local/opt/go/libexec
RUN export PATH=$PATH:$GOPATH/bin
RUN export PATH=$PATH:$GOROOT/bin

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/WannaChat

COPY ./ .

RUN glide install

RUN glide up

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]
