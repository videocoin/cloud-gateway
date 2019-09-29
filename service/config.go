package service

import "github.com/sirupsen/logrus"

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	Addr            string `default:"0.0.0.0:8080"`
	UsersRpcAddr    string `default:"0.0.0.0:5000"`
	StreamsRpcAddr  string `default:"0.0.0.0:5002"`
	ProfilesRpcAddr string `default:"0.0.0.0:5004"`

	Logger *logrus.Entry `envconfig:"-"`
}
