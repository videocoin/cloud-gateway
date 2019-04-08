package service

import "github.com/sirupsen/logrus"

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	Addr            string `default:"127.0.0.1:8080"`
	UsersRpcAddr    string `default:"127.0.0.1:5000" envconfig:"VC_USERS_RPCADDR"`
	AccountsRpcAddr string `default:"127.0.0.1:6000" envconfig:"VC_ACCOUNTS_RPCADDR"`
	StreamsRpcAddr  string `default:"127.0.0.1:2000" envconfig:"VC_STREAMS_RPCADDR"`

	Logger *logrus.Entry `envconfig:"-"`
}
