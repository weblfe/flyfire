package data

import (
		"context"
		"github.com/go-kratos/kratos/v2/registry"
		"github.com/go-redis/redis/v8"
		"github.com/weblfe/flyfire/app/account/service/internal/conf"
		"github.com/weblfe/flyfire/pkg/cache"
		"github.com/weblfe/flyfire/pkg/orm"
		"xorm.io/xorm"

		"github.com/go-kratos/kratos/v2/log"
		"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
		NewData,NewAccountRepo,
		NewDbClient,NewRedisClient,
		NewDiscovery,NewRegistrar,
		)

// Data .
type Data struct {
	dbClient orm.Conn
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

func NewDiscovery(conf *conf.Registry) registry.Discovery {
		return nil
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
		return nil
}


func NewCacheClient(c *conf.Data) cache.Cache {
		return cache.New()
}

// NewData .
func NewData(c *conf.Data, client orm.Conn, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		dbClient: client,
	}, cleanup, nil
}

func (d *Data) GetDb(ctx context.Context) *xorm.Session {
	return d.dbClient.Context(ctx)
}
