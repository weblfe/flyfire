package orm

import (
	"fmt"
	"githu.com/weblfe/flyfire/pkg/env"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"xorm.io/xorm"
)

type Conn = xorm.EngineInterface

type Option struct {
	source string
	driver string
}

type Options func(*Option)

const (
	DriverMysql = "mysql"
)

func (o *Option) GetDriver() string {
	switch strings.ToLower(o.driver) {
	case DriverMysql:
		return DriverMysql
	default:
		return DriverMysql
	}
}

func (o *Option) GetSource() string {
	if o.source == "" {
		o.source = env.GetOr("APP_DB_SOURCE", "root:@root@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8mb4&loc=Local")
	}
	if !strings.HasPrefix(o.source, o.GetDriverSchema()) {
		return fmt.Sprintf("%s%s", o.GetDriverSchema(), o.source)
	}
	return o.source
}

func (o *Option) GetDriverSchema() string {
	return fmt.Sprintf("%s://", o.GetDriver())
}

func (o *Option) Option(opts ...Options) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithDriver(driver string) Options {
	return func(o *Option) {
		o.driver = driver
	}
}

func WithSource(source string) Options {
	return func(o *Option) {
		o.source = source
	}
}

func NewConnection(opts ...Options) (Conn, error) {
	var o = &Option{}
	o.Option(opts...)
	return xorm.NewEngine(o.GetDriver(), o.GetSource())
}

type PoolOption struct {
	driver   string
	sources  []string
	policies xorm.GroupPolicy
}

type PoolOptions func(*PoolOption)

func WithPoolDriver(driver string) PoolOptions {
	return func(option *PoolOption) {
		option.driver = driver
	}
}

func WithPoolSources(sources ...string) PoolOptions {
	return func(option *PoolOption) {
		for _, v := range sources {
			if v != "" {
				option.sources = append(option.sources, v)
			}
		}
	}
}

func WithPoolPolices(polices ...xorm.GroupPolicy) PoolOptions {
	return func(option *PoolOption) {
		if len(polices) <= 0 || polices[0] == nil {
			return
		}
		option.policies = polices[0]
	}
}

func (o *PoolOption) GetSources() []string {
	if len(o.sources) <= 0 {
		o.sources = []string{
				env.GetOr("APP_DB_SOURCE","root:@root@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8mb4&loc=Local"),
		}
	}
	var sources []string
	for _, v := range o.sources {
		if !strings.HasPrefix(v, o.GetDriverSchema()) {
			sources = append(sources, fmt.Sprintf("%s%s", o.GetDriverSchema(), v))
		}
	}
	return sources
}

func (o *PoolOption) GetDriverSchema() string {
	return fmt.Sprintf("%s://", o.GetDriver())
}

func (o *PoolOption) GetDriver() string {
	switch strings.ToLower(o.driver) {
	case DriverMysql:
		return DriverMysql
	default:
		return DriverMysql
	}
}

func (o *PoolOption) Option(opts ...PoolOptions) {
	for _, opt := range opts {
		opt(o)
	}
}

func NewDb(opts ...PoolOptions) (Conn, error) {
	var o = &PoolOption{}
	o.Option(opts...)
	return xorm.NewEngineGroup(o.GetDriver(), o.GetSources(), o.policies)
}
