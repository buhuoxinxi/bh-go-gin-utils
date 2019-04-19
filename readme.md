# bh-go-gin-utils

捕获科技 gin 工具

## config

see ./testdata/config.go

## dev environment

go version

> go version go1.12 darwin/amd64 & go module enable

## test

go run server

```bash

go run ./testdata/executable/server.go
# INFO[0000] server addr :8099, isHTTPS(true)  


# ssl enable
curl -k https://127.0.0.1:8099/ping
# {"message":"pong","url":"/ping"}
curl -k https://127.0.0.1:8099/v1/ping
# {"message":"pong","url":"/v1/ping"}


# ssl unable
curl -k http://127.0.0.1:8099/ping
# {"message":"pong","url":"/ping"}
curl -k http://127.0.0.1:8099/v1/ping
# {"message":"pong","url":"/v1/ping"}

```
