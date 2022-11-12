package cache_test

import (
		"context"
		"github.com/stretchr/testify/assert"
		"github.com/weblfe/flyfire/pkg/cache"
		"github.com/weblfe/flyfire/pkg/env"
		"testing"
)

func init() {
	env.AutoLoad()
}

func TestNewCache(t *testing.T) {
	var (
		at     = assert.New(t)
		client = cache.New()
	)
	at.NotNil(client.Store(), "构建redis 客户端失败")
}

func TestCacheImpl_Get(t *testing.T) {
	var (
		at     = assert.New(t)
		client = cache.New()
		ctx = context.Background()
		value  = `test`
		key    = `test_one`
		v      string
	)
	at.Nil(client.Set(ctx,key, value), "添加缓存失败")
	at.Nil(client.Get(ctx,key, &v), "获取失败")
	at.Equal(value, v, "获取数据失败")
	at.Nil(client.Delete(ctx,key), "删除失败")
}
