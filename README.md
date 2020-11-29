# WannaChat

WannaChat - the ultimate collaboration platform for developers by developers

Our goal is to build the greatest collaboration platform for developer, are you ready to ditch Discord? Slack? Teams?

documentation: https://documenter.getpostman.com/view/3827865/Szf24qG7?version=latest

## Requirement:

- install go

## Testing and run:

```
// install deps
$ go mod tidy

// use go start server
$ go run main.go

open localhost:8080

// use air start server
$ air

open localhost:3000

// run test case
$ cd src/test
$ go test -v

// format code
$ go fmt
```

## Docker:

```
// build images and start container in one line
docker-compose up -d --build

// go inside container
docker exec -it <containerId> /bin/bash

// check container logs
docker logs <containerId>

// remove and stop container
docker-compose down
```

open localhost:8080

## Contributing

Please refer to [CONTRIBUTING.md](https://github.com/yeukfei02/WannaChat/blob/master/CONTRIBUTING.md)
