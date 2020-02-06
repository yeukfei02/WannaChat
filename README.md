# WannaChat
WannaChat - the ultimate collaboration platform for developers by developers

# Aim
Our goal is to build the greatest collaboration platform for developer, are you ready to ditch Discord? Slack? Teams?

### Requirement:
 - install go
 - install dep
 - install gin (https://github.com/codegangsta/gin)

### Testing and run:

clone the project to your GOPATH/src

check table in postgres
```
$ psql -h ec2-184-73-197-211.compute-1.amazonaws.com -U asgevskdmcckiw -d d7paprak57t9n5
$ 0e61a105e8512dc32ef63891155ef9dedb71237faf76cf5884b1a6797159bc44
```

```
$ dep ensure

// use go start server
$ go run main.go

open localhost:8080

// use gin start server
$ gin run main.go

open localhost:3000
```

### Docker:

- Dockerfile

build images and start container
```
docker build -t <username>/wanna-chat .
docker run -p 8080:8080 -d <username>/wanna-chat
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
