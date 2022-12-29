# mRPC

> A collection of wrappers for gRPC Lots of references to zRPC


# Run

```go
type depositServer struct{}

func newDepositServer() *depositServer {
	return &depositServer{}
}

func (s *depositServer) Deposit(ctx context.Context, in *mock.DepositRequest) (*mock.DepositResponse, error) {
	fmt.Printf("mrpc client in %+v \n", in)
	return &mock.DepositResponse{
		Ok: true,
	}, nil
}

func TestMrpc(t *testing.T) {
	// start server
	srv := newDepositServer()

	sc := RpcServerConf{
		Addr:    "localhost:8081",
		Timeout: 60,
	}
	s := MustNewServer(sc, func(server *grpc.Server) {
		mock.RegisterDepositServiceServer(server, srv)
	})

	defer s.Stop()

	fmt.Printf("starting rpc server at %s...\n", sc.Addr)

	go s.Start()

	// start client
	cc := RpcClientConf{
		Target:  "localhost:8081",
		Timeout: 60,
	}
	deposit := mock.NewDepositServiceClient(MustNewClient(cc).Conn())

	fmt.Printf("connection mrpc server %s... \n", cc.Target)

	resp, err := deposit.Deposit(context.Background(), &mock.DepositRequest{
		Amount: 10,
	})
	if err != nil {
		fmt.Printf("mrpc deposit err %v \n", err)
		return
	}

	fmt.Printf("mrpc deposit succ %+v \n", resp)
}
```

# Examples

- [mall-go](https://github.com/cexll/mall-go)


# License
Apache License Version 2.0, http://www.apache.org/licenses/