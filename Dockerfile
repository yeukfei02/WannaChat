FROM golang:1.16-alpine

RUN mkdir -p /app

WORKDIR /app

COPY ./ .

RUN go mod tidy

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]
