package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/weblfe/flyfire/app/account/service/internal/conf"
	"github.com/weblfe/flyfire/pkg/cache"
	"github.com/weblfe/flyfire/pkg/orm"
	"github.com/weblfe/flyfire/pkg/registry/etcd"
	"github.com/weblfe/flyfire/pkg/registry/nacos"
	"xorm.io/xorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewAccountRepo,
	NewDbClient, NewRedisClient,
	NewDiscovery, NewRegistrar,
	NewCacheClient,
)

// Data .
type Data struct {
	dbClient  orm.Conn
	showDebug bool
}

func NewDbClient(c *conf.Data) (orm.Conn, error) {
	return orm.NewDb(
		orm.WithPoolDriver(c.Database.Driver),
		orm.WithPoolSources(c.Database.Source),
	)
}

func NewRedisClient(c *conf.Redis) (*redis.Client, error) {
	return cache.NewClient(
		cache.WithAuth(c.Auth),
		cache.AddressSource(c.Addr),
		cache.WithNetwork(c.Network),
	)
}

func NewDiscovery(conf *conf.Registry) (registry.Discovery, error) {
	if conf.Nacos != nil {
		var servers []nacos.ServerConfig
		for _, s := range conf.Nacos.Servers {
			servers = append(servers, s)
		}
		return nacos.NewNaCosDiscoveryWiths(nacos.WithServers(servers...)), nil
	}
	if conf.Etcd != nil {
		return etcd.NewEtcdDiscoveryWiths(), nil
	}
	return nil, errors.New("discovery undefined")
}

func NewRegistrar(conf *conf.Registry) (registry.Registrar, error) {
	if conf.Etcd != nil {
		return etcd.NewEtcdRegistrarWiths(), nil
	}
	if conf.Nacos != nil {
		cc := conf.Nacos
		var arr []nacos.SourceConfig
		for _, v := range cc.ConfigSources {
			arr = append(arr, v)
		}
		return nacos.NewNaCosRegistrarWiths(
			nacos.WithSources(nacos.NewSources(arr...)...),
			nacos.WithGlobal(nacos.NewGlobal(conf.GetAppName())),
		), nil
	}
	return nil, errors.New(`registrar undefined`)
}

func NewCacheClient(c *conf.Data) cache.Cache {
	var (
		cnf        = c.GetCache()
		properties = cnf.GetProperties()
	)
	switch cnf.GetDriver() {
	case `redis`:
		opts := cache.Address(
			properties["host"],
			properties["port"])
		return cache.New(opts)
	case `local`:
		return cache.New()
	}
	return cache.New()
}

// NewData .
func NewData(c *conf.Data, client orm.Conn, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dbConf := c.GetDatabase()
	return &Data{
		dbClient:  client,
		showDebug: dbConf.GetShowDebug(),
	}, cleanup, nil
}

func (d *Data) GetDb(ctx context.Context) *xorm.Session {
	if !d.showDebug {
		return d.dbClient.Context(ctx)
	}
	return d.dbClient.Context(ctx).
		MustLogSQL(true)
}
