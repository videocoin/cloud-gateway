package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/cloud-gateway/service"
	"github.com/videocoin/cloud-pkg/logger"
)

var (
	ServiceName string = "gateway"
	Version     string = "dev"
)

func main() {
	logger.Init(ServiceName, Version)  //nolint

	log := logrus.NewEntry(logrus.New())
	log = logrus.WithFields(logrus.Fields{
		"service": ServiceName,
		"version": Version,
	})

	cfg := &service.Config{
		Name:    ServiceName,
		Version: Version,
	}

	err := envconfig.Process(ServiceName, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg.Logger = log

	gw, err := service.NewRPCGateway(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	signals := make(chan os.Signal, 1)
	exit := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		log.Infof("recieved signal %s", sig)
		exit <- true
	}()

	log.Info("starting")
	go log.Error(gw.Start())

	<-exit

	log.Info("stopping")
	err = gw.Stop()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("stopped")
}
