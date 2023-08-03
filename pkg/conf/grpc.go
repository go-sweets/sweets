package conf

import (
	"github.com/zeromicro/go-zero/core/prometheus"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/zrpc"
)

type GrpcConf struct {
	Addr       string `json:"Addr,default=localhost:8081"`
	Mode       string `json:",default=pro,options=dev|test|rt|pre|pro"`
	MetricsUrl string `json:",optional"`
	// Deprecated: please use DevServer
	Prometheus   prometheus.Config `json:",optional"`
	Telemetry    trace.Config      `json:",optional"`
	Auth         bool              `json:",optional"`
	Timeout      int64             `json:",default=2000"`
	CpuThreshold int64             `json:",default=900,range=[0:1000]"`
	// grpc health check switch
	Health      bool `json:",default=true"`
	Middlewares zrpc.ServerMiddlewaresConf
}
