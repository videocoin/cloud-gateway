package service

import (
	"context"
	"net/http"

	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/sirupsen/logrus"
	billingv1 "github.com/videocoin/cloud-api/billing/v1"
	dispatcherv1 "github.com/videocoin/cloud-api/dispatcher/v1"
	msv1 "github.com/videocoin/cloud-api/mediaserver/v1"
	minersv1 "github.com/videocoin/cloud-api/miners/v1"
	profilemanagerv1 "github.com/videocoin/cloud-api/profiles/manager/v1"
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

type RPCGateway struct {
	cfg    *Config
	logger *logrus.Entry
	e      *echo.Echo
	gw     *runtime.ServeMux
}

func NewRPCGateway(cfg *Config) (*RPCGateway, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	gw := &RPCGateway{
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

	err := profilemanagerv1.RegisterProfileManagerServiceHandlerFromEndpoint(ctx, mux, gw.cfg.ProfileManagerRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = usersv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, gw.cfg.UsersRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = streamsv1.RegisterStreamServiceHandlerFromEndpoint(ctx, mux, gw.cfg.StreamsRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = profilesv1.RegisterProfilesServiceHandlerFromEndpoint(ctx, mux, gw.cfg.ProfilesRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = minersv1.RegisterMinersServiceHandlerFromEndpoint(ctx, mux, gw.cfg.MinersRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = msv1.RegisterMediaServerServiceHandlerFromEndpoint(ctx, mux, gw.cfg.MediaServerRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = billingv1.RegisterBillingServiceHandlerFromEndpoint(ctx, mux, gw.cfg.BillingRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	err = dispatcherv1.RegisterDispatcherServiceHandlerFromEndpoint(ctx, mux, gw.cfg.DispatcherRPCAddr, opts)
	if err != nil {
		return nil, err
	}

	gw.gw = mux

	gw.route()

	return gw, nil
}

func (gw *RPCGateway) Start(errCh chan error) {
	gw.logger.Infof("starting gateway on %s", gw.cfg.Addr)
	go func() {
		errCh <- gw.e.Start(gw.cfg.Addr)
	}()

}

func (gw *RPCGateway) Stop() error {
	return nil
}

func (gw *RPCGateway) route() {
	gw.e.Use(middleware.CORS())

	gw.e.GET("/healthz", gw.health)
	gw.e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	gw.e.Any("/api/v1/*", echo.WrapHandler(gw.gw))
}

func (gw *RPCGateway) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"alive": "OK"})
}
