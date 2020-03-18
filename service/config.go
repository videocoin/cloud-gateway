package service

import "github.com/sirupsen/logrus"

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	Addr                  string `default:"0.0.0.0:8080"`
	UsersRPCAddr          string `default:"0.0.0.0:5000"`
	StreamsRPCAddr        string `default:"0.0.0.0:5002"`
	ProfilesRPCAddr       string `default:"0.0.0.0:5004"`
	ProfileManagerRPCAddr string `default:"0.0.0.0:5084"`
	MinersRPCAddr         string `default:"0.0.0.0:5011"`
	MediaServerRPCAddr    string `default:"0.0.0.0:5090"`
	BillingRPCAddr        string `default:"0.0.0.0:5020"`

	Logger *logrus.Entry `envconfig:"-"`
}
