# go-mixplus
A cloud-native Go microservices framework with cli tool for productivity.

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
├── Dockerfile
├── LICENSE
├── Makefile
├── api
│   ├── buf.lock
│   ├── buf.yaml
│   ├── hello
│   │   └── v1
│   │       ├── hello.pb.go
│   │       ├── hello.pb.gw.go
│   │       ├── hello.pb.validate.go
│   │       ├── hello.proto
│   │       └── hello_grpc.pb.go
│   └── openapi.yaml
├── buf.gen.yaml
├── buf.work.yaml
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── dbconfig.yml
├── etc
│   ├── config.yaml
│   └── gen.yml
├── go.mod
├── go.sum
└── internal
    ├── README.md
    ├── boundedcontexts
    │   └── hello
    │       ├── application
    │       │   └── handlers
    │       │       └── hello_grpc_handler.go
    │       ├── domain
    │       │   ├── entities
    │       │   │   └── user.go
    │       │   └── repositories
    │       │       └── hello_responsitories.go
    │       ├── infrastructure
    │       │   └── repositories
    │       │       └── hello_responsitories.go
    │       └── wire.go
    ├── config
    │   └── config.go
    ├── db
    │   ├── migration.go
    │   └── migrations
    │       ├── 202307101018-migrate-user-table.go
    │       ├── 20230718215456-create_data_migrations_table.sql
    │       └── db.go
    ├── server
    │   ├── grpc.go
    │   ├── http.go
    │   └── server.go
    ├── service
    │   ├── hello.go
    │   └── service.go
    └── svc
        └── serviceContext.go
```

the generated code can be run directly:

```
go mod tidy
make run
```

by default, it’s listening on port 8080, while it can be changed in the configuration file.

you can check it by curl:

```
curl -i 'http://localhost:8080/v1/demo/hello?id=1'
```
the response looks like below:

```
HTTP/1.1 200 OK
Content-Length: 32
Connection: keep-alive
Content-Type: application/json
Date: Wed, 12 Apr 2023 07:22:12 GMT
Keep-Alive: timeout=4
Proxy-Connection: keep-alive

{"id":"1","message":"Hello 1 !"}%
```

# Upgrade

```
go install github.com/mix-plus/go-mixplus/tools/mpctl@latest
```

# LICENSE
Apache License Version 2.0, http://www.apache.org/licenses/
