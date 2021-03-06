package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "deis-riak"
)

type Config struct {
	RiakMaster            bool   `envconfig:"RIAK_MASTER" default:"false"`
	ClusterServerHTTPPort int    `envconfig:"CLUSTER_SERVER_HTTP_PORT" default:"8080"`
	ClusterServerHost     string `envconfig:"DEIS_RIAK_CLUSTER_SERVICE_HOST" required:"true"`
	ClusterServerPort     int    `envconfig:"DEIS_RIAK_CLUSTER_SERVICE_PORT" required:"true"`
}

func Get() (*Config, error) {
	var ret Config
	if err := envconfig.Process(AppName, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
