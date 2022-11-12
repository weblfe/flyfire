package server

import (
		"errors"
		"github.com/go-kratos/kratos/v2/config"
		"github.com/go-kratos/kratos/v2/config/file"
		"github.com/go-kratos/kratos/v2/log"
		"github.com/weblfe/flyfire/app/account/service/internal/conf"
		"github.com/weblfe/flyfire/pkg/registry/nacos"
)

type SourceArgs struct {
	ConfDir string
	AppName string
	DataId  string
	Group   string
	log     log.Logger
}

// NewConfigure 构建用户配置
func NewConfigure(args *SourceArgs) (*conf.Bootstrap, config.Config, error) {
	var (
		err  error
		cfg  config.Config
		boot = new(conf.Bootstrap)
	)
	if cfg, err = NewConfigureSource(args); err != nil {
		return boot, cfg, err
	}
	if err = cfg.Scan(boot); err != nil {
		return boot, cfg, err
	}
	return boot, cfg, nil
}

func NewConfigureSource(source *SourceArgs) (config.Config, error) {
	var (
		err error
		cfg = config.New(
			config.WithSource(
				file.NewSource(source.ConfDir),
			),
		)
	)
	defer cfg.Close()

	if err = cfg.Load(); err != nil {
		return cfg, err
	}
	var rc = new(conf.Registry)
	if err = cfg.Scan(rc); err != nil {
		return cfg, err
	}

	if rc.Nacos != nil {
		return NewNacosConfig(rc, source)
	}
	if rc.Etcd != nil {
		return NewEtcdConfig(rc, source)
	}
	return cfg, nil
}

func NewEtcdConfig(rc *conf.Registry, source *SourceArgs) (cfg config.Config, err error) {
	return
}

func NewNacosConfig(rc *conf.Registry, source *SourceArgs) (cfg config.Config, err error) {
	var (
		logger   = source.log
		nacosCfg = rc.Nacos
		opts     *nacos.Config
	)
	if nacosCfg.DisableConfigClient {
		return
	}
	if rc.AppName == "" {
		rc.AppName = source.AppName
	}
	if len(nacosCfg.ConfigSources) < 1 {
		nacosCfg.ConfigSources = append(nacosCfg.ConfigSources,
			&conf.NacosConfigSources{DataId: source.DataId, Group: source.Group})
	}
	for _, s := range nacosCfg.ConfigSources {
		if s.DataId == "" {
			s.DataId = source.DataId
		}
		if s.Group == "" {
			s.Group = source.Group
		}
		_ = logger.Log(log.LevelInfo,
			"msg", "read config from nacos",
			"group", s.Group,
			"data_id", s.DataId,
			"namespace", rc.Nacos.Client.Namespace,
		)
	}

	if opts, err = toNacosCfg(rc); err != nil {
		return cfg, err
	}

	if cfg, err = nacos.FetchConfig(opts); err != nil {
		return cfg, err
	}
	return
}

func toNacosCfg(conf *conf.Registry) (*nacos.Config, error) {
	if conf.GetNacos() == nil {
		return &nacos.Config{}, errors.New("missing nacos config")
	}
	var serCfgArr []nacos.ServerConfig
	for _, s := range conf.Nacos.Servers {
		serCfgArr = append(serCfgArr, &nacos.Server{
			Port:    s.Port,
			Address: s.Address,
		})
	}
	var sourceArr []nacos.SourceConfig
	for _, s := range conf.Nacos.ConfigSources {
		sourceArr = append(sourceArr, &nacos.Source{Group: s.Group, DataId: s.DataId})
	}
	var (
		globalCfg = &nacos.Global{
			AppName: conf.GetAppName(),
		}
		cntCfg nacos.ClientConfig
	)
	if c := conf.Nacos.Client; c != nil {
		cntCfg = &nacos.ClientOption{
			Username:            c.Username,
			Password:            c.Password,
			CacheDir:            c.CacheDir,
			LogDir:              c.LogDir,
			LogLevel:            c.LogLevel,
			Timeout:             c.Timeout,
			Namespace:           c.Namespace,
			MaxBackups:          c.MaxBackups,
			MaxAge:              c.RollingMaxAge,
			MaxSize:             c.RollingMaxSize,
			UpdateThreadNum:     int(c.UpdateThreadNum),
			NotLoadCacheAtStart: c.GetNotLoadCacheAtStart(),
		}
	}
	return &nacos.Config{
		CntCfg:    cntCfg,
		SerCfgArr: serCfgArr,
		GlbCfg:    globalCfg,
		SrcCfgArr: sourceArr,
	}, nil
}

// Watcher 监听配置修改
func Watcher(cfg config.Config, conf *conf.Server, logger log.Logger) {
	var helper = log.NewHelper(logger)
	if err := cfg.Watch("*", restart(conf,helper)); err != nil {
		helper.Errorw("Watcher.Error", err)
	}
	helper.Info("service restart by configure update signal")
}

// 重启服务
func restart(conf *conf.Server, logHelper *log.Helper) func(s string, value config.Value) {
	return func(s string, value config.Value) {
		logHelper.Infow("key", s, "value", value.Load())

	}
}
