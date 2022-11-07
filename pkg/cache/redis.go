package cache

import (
	"context"
	"githu.com/weblfe/flyfire/pkg/codec"
	"githu.com/weblfe/flyfire/pkg/env"
	red "github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const netTcp = `tcp`

type Cache interface {
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string, typeVal interface{}) error
	SetDefault(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, expire ...time.Duration) error
}

type RedisStore interface {
	Store() *red.Client
	SetStore(client *red.Client)
}

type Option func(*cacheOptions)

type cacheImpl struct {
	conn    *red.Client
	locker  sync.Locker
	options *cacheOptions
}

type cacheOptions struct {
	db                 int
	network            string
	addr               string
	username           string
	password           string
	maxRetries         int
	poolSize           int
	minIdleConns       int
	defaultExpire      time.Duration
	maxConnAge         time.Duration
	minRetryBackoff    time.Duration
	maxRetryBackoff    time.Duration
	dialTimeout        time.Duration
	readTimeout        time.Duration
	writeTimeout       time.Duration
	poolTimeout        time.Duration
	idleTimeout        time.Duration
	idleCheckFrequency time.Duration
	codec              codec.Codec
}

func New(options ...Option) Cache {
	var impl = new(cacheImpl)
	impl.init()
	impl.Apply(options...)
	return impl
}

func NewClient(opts ...Option) (*red.Client, error) {
	var (
		opt    = newDefaultOptions().Apply(opts...)
		client = red.NewClient(opt.Option())
	)
	_, err := client.Ping(context.Background()).Result()
	return client, err
}

func newDefaultOptions() *cacheOptions {
	var opts = new(cacheOptions)
	opts.network = netTcp
	opts.defaultExpire = 3 * time.Minute
	opts.codec, _ = codec.GetCodec("json")
	opts.username = env.GetOr(`REDIS_USERNAME`)
	opts.password = env.GetOr(`REDIS_PASSWORD`)
	opts.db = env.GetIntOr(`REDIS_DB`, 0)
	opts.addr = env.GetOr(`REDIS_HOST`, `127.0.0.1:6379`)
	opts.minIdleConns = env.GetIntOr(`REDIS_MIN_IDLE_CONNS`, 3)
	return opts
}

func (c *cacheImpl) init() {
	c.locker = &sync.RWMutex{}
	c.options = newDefaultOptions()
}

func (c *cacheImpl) SetConn(r *red.Client) Cache {
	if r == nil || c.conn == nil {
		return c
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	c.conn = r
	return c
}

func (c *cacheImpl) Apply(options ...Option) {
	for _, o := range options {
		o(c.options)
	}
	return
}

func (c *cacheImpl) Get(ctx context.Context, key string, result interface{}) error {
	value, err := c.Conn().Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return c.getCodec().Decode([]byte(value), result)
}

func (c *cacheImpl) getCodec() codec.Codec {
	return c.options.codec
}

func (c *cacheImpl) Set(ctx context.Context, key string, value interface{}, expire ...time.Duration) error {
	expire = append(expire, 0)
	var (
		exp       = expire[0]
		data, err = c.getCodec().Encode(value)
	)
	if exp <= 0 {
		exp = -1
	}
	if _, err = c.Conn().Set(ctx, key, data, exp).Result(); err != nil {
		return err
	}
	return nil
}

func (c *cacheImpl) SetDefault(ctx context.Context, key string, value interface{}) error {
	return c.Set(ctx, key, value, c.options.defaultExpire)
}

func (c *cacheImpl) Delete(ctx context.Context, key string) error {
	if _, err := c.Conn().Del(ctx, key).Result(); err != nil {
		return err
	}
	return nil
}

func (c *cacheImpl) Store() *red.Client {
	return c.Conn()
}

func (c *cacheImpl) SetStore(client *red.Client) {
	if c.conn == nil {
		c.conn = client
	}
	return
}

func (c *cacheImpl) Conn() *red.Client {
	if c.conn != nil {
		return c.conn
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	c.conn = red.NewClient(c.getRedOptions())
	return c.conn
}

func (c *cacheImpl) getRedOptions() *red.Options {
	return c.options.Option()
}

func (o *cacheOptions) Option() *red.Options {
	return &red.Options{
		DB:                 o.db,
		Addr:               o.addr,
		Network:            o.network,
		Username:           o.username,
		Password:           o.password,
		MaxRetries:         o.maxRetries,
		MinRetryBackoff:    o.minRetryBackoff,
		MaxRetryBackoff:    o.maxRetryBackoff,
		DialTimeout:        o.dialTimeout,
		ReadTimeout:        o.readTimeout,
		WriteTimeout:       o.writeTimeout,
		PoolSize:           o.poolSize,
		MinIdleConns:       o.minIdleConns,
		MaxConnAge:         o.maxConnAge,
		PoolTimeout:        o.poolTimeout,
		IdleTimeout:        o.idleTimeout,
		IdleCheckFrequency: o.idleCheckFrequency,
	}
}

func (o *cacheOptions) Apply(opts ...Option) *cacheOptions {
	for _, op := range opts {
		op(o)
	}
	return o
}
