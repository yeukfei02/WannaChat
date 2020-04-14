# WannaChat

WannaChat - the ultimate collaboration platform for developers by developers

Our goal is to build the greatest collaboration platform for developer, are you ready to ditch Discord? Slack? Teams?

## Requirement:
 - install go
 - install dep
 - install gin (https://github.com/codegangsta/gin)

## Testing and run:

clone the project to your GOPATH/src

```
$ dep ensure

// use go start server
$ go run main.go

open localhost:8080

// use gin start server
$ gin run main.go

open localhost:3000
```

## Docker:

- Dockerfile

build images and start container
```
docker build -t <username>/wanna-chat:<tag> .
docker run -p 8080:8080 -d <username>/wanna-chat:<tag>
docker exec -it <containerId> /bin/bash
docker logs <containerId>
```

check images and container
```
docker images
docker ps
docker ps -a
```

open localhost:8080

- docker-compose.yml

build images and start container
```
docker-compose build
docker-compose up
```
build images and start container in one line
```
docker-compose up -d --build
```

stop container
```
docker-compose stop
```

open localhost:8080

## Contributing

Please refer to [CONTRIBUTING.md](https://github.com/yeukfei02/WannaChat/blob/master/CONTRIBUTING.md)
