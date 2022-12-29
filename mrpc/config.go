package mrpc

const (
	DevMode     = "dev"
	TestMode    = "test"
	RtMode      = "rt"
	PreMode     = "pre"
	ProMode     = "pro"
	ReleaseMode = "release"
)

type (
	RpcServerConf struct {
		ServiceConf `mapstructure:",squash"`
		Addr        string `json:"Addr"`
		Timeout     int64  `json:"Timeout"`
	}
	RpcClientConf struct {
		Target   string `json:"Target"`
		NonBlock bool   `json:"NonBlock"`
		Timeout  int64  `json:"Timeout"`
	}
	ServiceConf struct {
		Name       string `json:"Name"`
		Mode       string `json:"Mode"`
		MetricsUrl string `json:"MetricsUrl"`
		Prometheus string `json:"Prometheus"`
	}
)
