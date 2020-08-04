package service

import "github.com/sirupsen/logrus"

type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	Addr                  string `default:"0.0.0.0:8080"`
	UsersRPCAddr          string `default:"0.0.0.0:5000"`
	StreamsRPCAddr        string `default:"0.0.0.0:5002"`
	MinersRPCAddr         string `default:"0.0.0.0:5011"`
	MediaServerRPCAddr    string `default:"0.0.0.0:5090"`
	BillingRPCAddr        string `default:"0.0.0.0:5020"`
	DispatcherRPCAddr     string `default:"0.0.0.0:5008"`
}
