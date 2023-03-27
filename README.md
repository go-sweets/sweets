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
.
├── LICENSE
├── Makefile
├── README.md
├── api
│   ├── buf.lock
│   ├── buf.yaml
│   ├── hello
│   │   └── v1
│   │       ├── hello.pb.go
│   │       ├── hello.pb.gw.go
│   │       ├── hello.pb.validate.go
│   │       ├── hello.proto
│   │       └── hello_grpc.pb.go
│   └── openapi.yaml
├── buf.gen.yaml
├── buf.work.yaml
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── etc
│   └── config.yaml
├── go.mod
├── go.sum
└── internal
    ├── README.md
    ├── config
    │   └── config.go
    ├── server
    │   ├── grpc.go
    │   ├── http.go
    │   └── server.go
    ├── service
    │   ├── hello.go
    │   └── service.go
    └── svc
        └── serviceContext.go

11 directories, 27 files
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


# LICENSE
Apache License Version 2.0, http://www.apache.org/licenses/