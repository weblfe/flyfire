package main

import (
		"flag"
		"github.com/go-kratos/kratos/v2/config"
		"github.com/go-kratos/kratos/v2/registry"
		"github.com/weblfe/flyfire/app/account/service/internal/server"
		"net/url"
		"os"

		"github.com/weblfe/flyfire/app/account/service/internal/conf"

		"github.com/go-kratos/kratos/v2"
		"github.com/go-kratos/kratos/v2/log"
		"github.com/go-kratos/kratos/v2/middleware/tracing"
		"github.com/go-kratos/kratos/v2/transport/grpc"
		"github.com/go-kratos/kratos/v2/transport/http"

		_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "account.firefly.service"
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	confDir string

	id, _ = os.Hostname()
	// DefaultDataID default config source
	DefaultDataID = Name + ".yaml"
	// DefaultGroup default config source
	DefaultGroup = "DEFAULT_GROUP"
)

func init() {
	flag.StringVar(&confDir, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(
	rr registry.Registrar,
	endpoints []*url.URL,
	logger log.Logger,
	gs *grpc.Server,
	hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Endpoint(endpoints...),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	var (
		err    error
		cfg    config.Config
		bc     *conf.Bootstrap
		logger = log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
		args = &server.SourceArgs{
			AppName: Name,
			ConfDir: confDir,
			DataId:  DefaultDataID,
			Group:   DefaultGroup,
		}
		rr = new(conf.Registry)
	)

	if bc, cfg, err = server.NewConfigure(args); err != nil {
		panic(err)
	}
	// 服务配置监听
	server.Watcher(cfg, bc.Server, logger)

	app, cleanup, err := wireApp(bc.Endpoints,rr ,bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}

}
