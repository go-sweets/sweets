# mix-plus-cli
> processing


mix-plus-cli init mix-plus go skeleton


# Install
```go
go get github.com/mix-plus/mpcli
```

# Run
```
// init grpc
mpcli new mrpc
// init api 
mpcli new api
// init queue
mpcli new mqueue
```

## 下载逻辑
利用 go mod 下载包到本地GOPATH `\pkg\mod\cache\download` 目录然后解压到`cli`运行目录替换命名空间

# License
Apache License Version 2.0, http://www.apache.org/licenses/

