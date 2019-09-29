package service

import (
	"context"
	"net/http"

	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	profilesv1 "github.com/videocoin/cloud-api/profiles/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
	usersv1 "github.com/videocoin/cloud-api/users/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"google.golang.org/grpc/metadata"
)

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	var (
		otHeaders = []string{
			"x-request-id",
			"x-b3-traceid",
			"x-b3-spanid",
			"x-b3-parentspanid",
			"x-b3-sampled",
			"x-b3-flags",
			"x-ot-span-context"}
	)
	var pairs []string

	for _, h := range otHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

type annotator func(context.Context, *http.Request) metadata.MD

func chainGrpcAnnotators(annotators ...annotator) annotator {
	return func(c context.Context, r *http.Request) metadata.MD {
		mds := []metadata.MD{}
		for _, a := range annotators {
			mds = append(mds, a(c, r))
		}
		return metadata.Join(mds...)
	}
}

type RpcGateway struct {
	cfg    *Config
	logger *logrus.Entry
	e      *echo.Echo
	gw     *runtime.ServeMux
}

func NewRpcGateway(cfg *Config) (*RpcGateway, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

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

	annotators := []annotator{injectHeadersIntoMetadata}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler),
		runtime.WithMetadata(chainGrpcAnnotators(annotators...)),
		grpcutil.WithProtoHTTPErrorHandler(),
	)
	opts := grpcutil.DefaultClientDialOpts(cfg.Logger)

	err := usersv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, gw.cfg.UsersRpcAddr, opts)
	if err != nil {
		return nil, err
	}

	err = streamsv1.RegisterStreamServiceHandlerFromEndpoint(ctx, mux, gw.cfg.StreamsRpcAddr, opts)
	if err != nil {
		return nil, err
	}

	err = profilesv1.RegisterProfilesServiceHandlerFromEndpoint(ctx, mux, gw.cfg.ProfilesRpcAddr, opts)

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
