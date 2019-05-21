package service

import (
	"context"
	"net/http"

	accountsv1 "github.com/VideoCoin/cloud-api/accounts/v1"
	pipelinesv1 "github.com/VideoCoin/cloud-api/pipelines/v1"
	usersv1 "github.com/VideoCoin/cloud-api/users/v1"

	"github.com/VideoCoin/cloud-pkg/grpcutil"
	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type RpcGateway struct {
	cfg    *Config
	logger *logrus.Entry
	e      *echo.Echo
	gw     *runtime.ServeMux
}

func NewRpcGateway(cfg *Config) (*RpcGateway, error) {
	e := echo.New()
	e.HideBanner = true

	gw := &RpcGateway{
		cfg:    cfg,
		logger: cfg.Logger,
		e:      e,
	}

	ctx := context.Background()
	marshaler := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler),
		grpcutil.WithProtoHTTPErrorHandler(),
	)
	opts := grpcutil.DefaultClientDialOpts(cfg.Logger)

	err := accountsv1.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, gw.cfg.AccountsRpcAddr, opts)
	if err != nil {
		return nil, err
	}

	err = usersv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, gw.cfg.UsersRpcAddr, opts)
	if err != nil {
		return nil, err
	}

	err = pipelinesv1.RegisterPipelineServiceHandlerFromEndpoint(ctx, mux, gw.cfg.PipelinesRpcAddr, opts)
	if err != nil {
		return nil, err
	}

	gw.gw = mux

	gw.route()

	return gw, nil
}

func (gw *RpcGateway) Start() error {
	gw.logger.Infof("starting gateway on %s", gw.cfg.Addr)
	return gw.e.Start(gw.cfg.Addr)
}

func (gw *RpcGateway) Stop() error {
	return nil
}

func (gw *RpcGateway) route() {
	gw.e.Use(middleware.CORS())

	gw.e.GET("/healthz", gw.health)
	gw.e.GET("/metrics", echo.WrapHandler(prometheus.Handler()))
	gw.e.Any("/api/v1/*", echo.WrapHandler(gw.gw))
}

func (gw *RpcGateway) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"alive": "OK"})
}
