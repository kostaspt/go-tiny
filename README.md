# Tiny
> Boilerplate for microservice written in Go.

### How to build docker image
```shell
$ docker build --build-arg VERSION=$(git describe --tags --always) --build-arg SQL_ADDR=sql:5432 --build-arg SERVER_PORT=4000 -t tiny:latest .
```