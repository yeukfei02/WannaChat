# WannaChat

WannaChat

documentation: <https://documenter.getpostman.com/view/3827865/Szf24qG7?version=latest>

url: <https://wanna-chat-api.herokuapp.com/>

## Requirement

- install go

## Testing and run

```zsh
// install deps
$ go mod tidy

// use go start server
$ go run main.go

// use air start server
$ air

open localhost:8080

// run test case
$ cd src/test
$ go test -v

// format code
$ go fmt
```

## Docker

```zsh
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
