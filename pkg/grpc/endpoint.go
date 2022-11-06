package grpc

import (
	"fmt"
	"strings"
)

const (
	DiscoverySchema   = "discovery:///"
	DefaultServiceTpl = "app.%s.service.grpc"
)

type Selector interface {
	Get(name string) string
}

type Option struct {
	template string
	schema   string
	cache    map[string]string
}

type Options func(*Option)

func GetEndPoint(selector Selector, service string) string {
	return selector.Get(service)
}

func Resolver(conf map[string]string, service string) string {
	return GetEndPoint(newSelector(conf, DiscoverySchema, DefaultServiceTpl), service)
}

func AutoResolver(service string, opts ...Options) string {
	return GetEndPoint(NewSelector(opts...), service)
}

type selectorImpl struct {
	schema   string
	template string
	cache    map[string]string
}

func (s *selectorImpl) Get(name string) string {
	var (
		defaultEndpoint string
		schema          = s.getSchema()
	)
	name = strings.ToLower(name)
	if strings.HasPrefix(name, schema) {
		defaultEndpoint = name
	} else {
		defaultEndpoint = s.resolver(schema, name)
	}
	if len(s.cache) == 0 {
		return defaultEndpoint
	}
	endpoint, ok := s.cache[name]
	if !ok {
		return defaultEndpoint
	}
	return endpoint
}

func (s *selectorImpl) getSchema() string {
	if s.schema == "" {
		return DiscoverySchema
	}
	return s.schema
}

func (s *selectorImpl) resolver(schema, name string) string {
	var template = s.template
	if template == "" {
		template = DefaultServiceTpl
	}
	if strings.HasPrefix(template, schema) {
		return fmt.Sprintf(template, name)
	}
	return fmt.Sprintf(template, schema, name)
}

func newSelector(serverCfg map[string]string, proto, template string) Selector {
	return &selectorImpl{
		schema:   proto,
		template: template,
		cache:    serverCfg,
	}
}

func NewSelector(opts ...Options) Selector {
	var impl = &selectorImpl{
		schema:   DiscoverySchema,
		template: DefaultServiceTpl,
		cache:    map[string]string{},
	}
	opt := &Option{}
	for _, o := range opts {
		o(opt)
	}
	return impl.Apply(opt)
}

func (s *selectorImpl) Apply(o *Option) Selector {
	if o.schema != "" && s.schema != o.schema {
		s.schema = o.schema
	}
	if o.template != "" && s.template != o.template {
		s.template = o.template
	}
	if len(o.cache) > 0 {
		for k, v := range o.cache {
			if v != "" && k != "" {
				s.cache[k] = v
			}
		}
	}
	return s
}
