package nacos

import (
	"errors"
	cfgNaCos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	nacosV2 "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	GlbCfg    GlobalConfig   `json:"global" yaml:"global"`
	SerCfgArr []ServerConfig `json:"servers" yaml:"servers"`
	CntCfg    ClientConfig   `json:"client_config" yaml:"client_config"`
	SrcCfgArr []SourceConfig `json:"sources" yaml:"sources"`
}

type GlobalConfig interface {
	GetAppName() string
}

type ServerConfig interface {
	GetAddress() string
	GetPort() uint64
}

type ClientConfig interface {
	RollingConfig

	GetNamespace() string
	GetTimeout() uint64
	GetUsername() string
	GetPassword() string
	GetCacheDir() string
	GetLogDir() string
	GetLogLevel() string
	GetUpdateThreadNum() int
	GetNotLoadCacheAtStart() bool
}

type RollingConfig interface {
	GetMaxAge() int64
	GetMaxSize() int64
	GetMaxBackups() int64
}

type SourceConfig interface {
	GetDataId() string
	GetGroup() string
}

type Option func(*Config)

func NewNaCosConfig(opts ...Option) *Config {
	var c = &Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Config) GetGlobalConfig() GlobalConfig {
	return c.GlbCfg
}

type Global struct {
	AppName string `json:"app_name" yaml:"app_name"`
}

func WithGlobal(g *Global) Option {
	return func(c *Config) {
		c.GlbCfg = g
	}
}

func NewGlobal(name string) *Global {
	return &Global{
		AppName: name,
	}
}

func NewSources[T any](sources ...T) []T {
	return sources
}

func (g *Global) GetAppName() string {
	return g.AppName
}

type Server struct {
	Port    uint64 `json:"port" yaml:"port"`
	Address string `json:"address" yaml:"address"`
}

func WithServers(serv ...ServerConfig) Option {
	return func(c *Config) {
		for _, v := range serv {
			if v == nil {
				continue
			}
			c.SerCfgArr = append(c.SerCfgArr, v)
		}
	}
}

func (s *Server) GetAddress() string {
	return s.Address
}

func (s *Server) GetPort() uint64 {
	return s.Port
}

type Source struct {
	Group  string `yaml:"group" json:"group"`
	DataId string `yaml:"data_id" json:"data_id"`
}

type ClientOption struct {
	Username            string `json:"username" yaml:"username"`
	Password            string `json:"password" yaml:"password"`
	CacheDir            string `json:"cache_dir" yaml:"cache_dir"`
	LogDir              string `json:"log_dir" yaml:"log_dir"`
	LogLevel            string `json:"log_level" yaml:"log_level"`
	Timeout             uint64 `json:"timeout" yaml:"timeout"`
	Namespace           string `json:"namespace" yaml:"namespace"`
	MaxBackups          int64  `json:"max_backups" yaml:"max_backups"`
	MaxAge              int64  `json:"rolling_max_age" yaml:"rolling_max_age"`
	MaxSize             int64  `json:"rolling_max_size" yaml:"rolling_max_size"`
	UpdateThreadNum     int    `json:"update_thread_num" yaml:"update_thread_num"`
	NotLoadCacheAtStart bool   `json:"not_load_cache_at_start" yaml:"not_load_cache_at_start"`
}

func (c *ClientOption) GetMaxAge() int64 {
	return c.MaxAge
}

func (c *ClientOption) GetMaxSize() int64 {
	return c.MaxSize
}

func (c *ClientOption) GetMaxBackups() int64 {
	return c.MaxBackups
}

func (c *ClientOption) GetNamespace() string {
	return c.Namespace
}

func (c *ClientOption) GetTimeout() uint64 {
	return c.Timeout
}

func (c *ClientOption) GetUsername() string {
	return c.Username
}

func (c *ClientOption) GetPassword() string {
	return c.Password
}

func (c *ClientOption) GetCacheDir() string {
	return c.CacheDir
}

func (c *ClientOption) GetLogDir() string {
	return c.LogDir
}

func (c *ClientOption) GetLogLevel() string {
	return c.LogLevel
}

func (c *ClientOption) GetNotLoadCacheAtStart() bool {
	return c.NotLoadCacheAtStart
}

func (c *ClientOption) GetUpdateThreadNum() int {
	return c.UpdateThreadNum
}

func WithClientOption(opt *ClientOption) Option {
	return func(c *Config) {
		c.CntCfg = opt
	}
}

func WithSources(source ...SourceConfig) Option {
	return func(c *Config) {
		for _, s := range source {
			if s == nil {
				continue
			}
			c.SrcCfgArr = append(c.SrcCfgArr, s)
		}
	}
}

func (s *Source) GetDataId() string {
	return s.DataId
}

func (s *Source) GetGroup() string {
	return s.Group
}

func NewNaCosDiscovery(conf *Config) registry.Discovery {
	client := NewNaCosNamingClient(conf)
	return nacosV2.New(client)
}

func NewNaCosRegistrar(conf *Config) registry.Registrar {
	client := NewNaCosNamingClient(conf)
	r := nacosV2.New(client)
	return r
}

func NewNaCosRegistrarWiths(opts ...Option) registry.Registrar {
	return NewNaCosRegistrar(NewNaCosConfig(opts...))
}

func NewNaCosDiscoveryWiths(opts ...Option) registry.Discovery {
	return NewNaCosDiscovery(NewNaCosConfig(opts...))
}

func NewNaCosClientParam(cfg *Config) vo.NacosClientParam {
	sc := make([]constant.ServerConfig, 0)

	cntCfg := cfg.CntCfg
	glbCfg := cfg.GlbCfg
	serCfgArr := cfg.SerCfgArr

	for _, s := range serCfgArr {
		sc = append(sc, *constant.NewServerConfig(s.GetAddress(), s.GetPort()))
	}
	logDir := cntCfg.GetLogDir()
	if logDir != "" && glbCfg.GetAppName() != "" {
		logDir = logDir + string(os.PathSeparator) + glbCfg.GetAppName()
	}

	cacheDir := cntCfg.GetCacheDir()
	if logDir != "" && glbCfg.GetAppName() != "" {
		cacheDir = cacheDir + string(os.PathSeparator) + glbCfg.GetAppName()
	}
	var options = CreateOptions(cntCfg, cacheDir, logDir)
	cc := constant.NewClientConfig(options...)
	cc.AppName = glbCfg.GetAppName()
	return vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
}

func NewNaCosConfigClient(conf *Config) config_client.IConfigClient {
	cfgClient, err := clients.NewConfigClient(
		NewNaCosClientParam(conf),
	)
	if err != nil {
		panic(err)
	}
	return cfgClient
}

func NewNaCosNamingClient(conf *Config) naming_client.INamingClient {
	client, err := clients.NewNamingClient(NewNaCosClientParam(conf))
	if err != nil {
		panic(err)
	}
	return client
}

func FetchConfig(conf *Config) (config.Config, error) {
	if len(conf.SrcCfgArr) < 1 {
		return nil, errors.New("missing config source")
	}
	confSources := make([]config.Source, 0)

	for _, cs := range conf.SrcCfgArr {
		confSources = append(confSources, cfgNaCos.NewConfigSource(
			NewNaCosConfigClient(conf),
			cfgNaCos.WithDataID(cs.GetDataId()),
			cfgNaCos.WithGroup(cs.GetGroup()),
		))
	}

	c := config.New(config.WithSource(confSources...))
	err := c.Load()
	return c, err
}

func CreateOptions(cntCfg ClientConfig, cacheDir, logDir string) (options []constant.ClientOption) {
	if cacheDir != "" {
		options = append(options, constant.WithCacheDir(cacheDir))
	}
	if logDir != "" {
		options = append(options, constant.WithLogDir(logDir))
	}
	if namespace := cntCfg.GetNamespace(); namespace != "" {
		options = append(options, constant.WithNamespaceId(namespace))
	}
	if timeout := cntCfg.GetTimeout(); timeout > 0 {
		options = append(options, constant.WithTimeoutMs(timeout))
	}
	if username := cntCfg.GetUsername(); username != "" {
		options = append(options, constant.WithUsername(username))
	}
	if pwd := cntCfg.GetPassword(); pwd != "" {
		options = append(options, constant.WithPassword(pwd))
	}
	if level := cntCfg.GetLogLevel(); level != "" {
		options = append(options, constant.WithLogLevel(level))
	}
	if max := cntCfg.GetMaxAge(); max > 0 {
		cfg := &lumberjack.Logger{
			LocalTime:  true,
			Compress:   true,
			MaxAge:     int(max),
			Filename:   emptyStrOr(cntCfg.GetLogDir(), getDefaultLogDir()),
			MaxBackups: emptyOIntOr(int(cntCfg.GetMaxBackups()), 3),
			MaxSize:    emptyOIntOr(int(cntCfg.GetMaxSize()), 512),
		}
		opt := constant.WithLogRollingConfig(cfg)
		options = append(options, opt)
	}
	if ok := cntCfg.GetNotLoadCacheAtStart(); ok {
		options = append(options, constant.WithNotLoadCacheAtStart(ok))
	}
	if num := cntCfg.GetUpdateThreadNum(); num > 0 {
		options = append(options, constant.WithUpdateThreadNum(int(num)))
	}
	return
}

func emptyStrOr(v, defaultVal string) string {
	if v == "" {
		return defaultVal
	}
	return v
}

func emptyOIntOr(n, defaultVal int) int {
	if n == 0 {
		return defaultVal
	}
	return defaultVal
}

func getDefaultLogDir() string {
	dir, _ := os.Getwd()
	if dir == "" {
		return "nacos.log"
	}
	return dir + string(os.PathSeparator) + "nacos.log"
}
