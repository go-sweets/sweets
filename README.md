# go-mixplus
A cloud-native Go microservices framework with cli tool for productivity.

# Installation
Run the following command under your project:

```
go get -u github.com/mix-plus/go-mixplus
```

# Upgrade

```
go install github.com/mix-plus/go-mixplus/tools/mpctl@latest
```

# Quick Start
1. install mpctl
```
# for Go 1.15 and earlier
GO111MODULE=on go get -u github.com/mix-plus/go-mixplus/tools/mpctl@latest

# for Go 1.16 and later
go install github.com/mix-plus/go-mixplus/tools/mpctl@latest

# generate 
mpctl new helloservice
```
the generated files look like

```
api
├── api
│   ├── hello.pb.go
│   └── hello.proto
├── cmd
│   └── server.go
├── etc
│   └── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── controllers
│   │   │   └── hello.go
│   │   └── routes.go
│   ├── logic
│   │   └── hello.go
│   └── svc
│       └── serviceContext.go
├── LICENSE
├── Makefile
└── README.md
```

the generated code can be run directly:

```
go mod tidy
make run
```

by default, it’s listening on port 8080, while it can be changed in the configuration file.

you can check it by curl:

```
curl -i 'http://localhost:8080/v1/dmeo/hello' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1
}'
```
the response looks like below:

```
HTTP/1.1 200 OK
Content-Length: 33
Connection: keep-alive
Content-Type: application/json
Date: Mon, 27 Mar 2023 06:08:32 GMT
Keep-Alive: timeout=4
Proxy-Connection: keep-alive

{"id":"1", "message":"Hello 1 !"}
```

# Documents
TODO

# LICENSE
Apache License Version 2.0, http://www.apache.org/licenses/