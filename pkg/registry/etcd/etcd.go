package etcd

import (
	"context"
	etcdV2 "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdClient "go.etcd.io/etcd/client/v3"
)

type Config struct {
	registryOpts []etcdV2.Option
	etcdOpts     []etcdClient.Option
}

type OptionEtcd func(config *Config)

func NewEtcdRegistrarWiths(opts ...OptionEtcd) registry.Registrar {
	return NewEtcdRegistrar(NewEtcdConfig(opts...))
}

func NewEtcdRegistrar(cfg *Config) registry.Registrar {
	return etcdV2.New(NewEtcdClient(cfg.etcdOpts...), cfg.registryOpts...)
}

func NewEtcdDiscoveryWiths(opts ...OptionEtcd) registry.Discovery {
	return NewEtcdDiscovery(NewEtcdConfig(opts...))
}

func NewEtcdDiscovery(cfg *Config) registry.Discovery {
	return etcdV2.New(NewEtcdClient(cfg.etcdOpts...), cfg.registryOpts...)
}

func NewEtcdClient(opts ...etcdClient.Option) *etcdClient.Client {
	return etcdClient.NewCtxClient(context.Background(), opts...)
}

func NewEtcdConfig(opts ...OptionEtcd) *Config {
	var cfg = &Config{}
	for _, o := range opts {
		o(cfg)
	}
	return cfg
}
