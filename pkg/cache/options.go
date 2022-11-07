package cache

import (
	"fmt"
	"githu.com/weblfe/flyfire/pkg/codec"
	"strconv"
	"time"
)

func User(user, password string) Option {
	return func(options *cacheOptions) {
		if user != "" && options.username == "" {
			options.username = user
		}
		if password != "" && options.password == "" {
			options.password = password
		}
	}
}

func Address(host, port string) Option {
	return func(options *cacheOptions) {
		if port == "" {
			port = "6379"
		}
		n, err := strconv.Atoi(port)
		if err != nil {
			return
		}
		if n <= 0 || n > 2^20 {
			return
		}
		if host != "" && options.addr == "" {
			options.addr = fmt.Sprintf("%s:%s", host, port)
		}
	}
}

func AddressSource(source string) Option {
	return func(options *cacheOptions) {
		if source != "" && options.addr == "" {
			options.addr = source
		}
	}
}

func WithAuth(auth string)Option  {
		return func(options *cacheOptions) {
				options.password = auth
		}
}

func WithNetwork(network string)Option  {
		return func(options *cacheOptions) {
				options.network  =network
		}
}

func ExpireDuration(d time.Duration) Option {
	return func(options *cacheOptions) {
		if d <= 0 {
			return
		}
		options.defaultExpire = d
	}
}

func Db(db int) Option {
	return func(options *cacheOptions) {
		if db <= 0 {
			return
		}
		options.db = db
	}
}

func MaxRetries(maxRetries int) Option {
	return func(options *cacheOptions) {
		options.maxRetries = maxRetries
	}
}

func PoolSize(poolSize int) Option {
	return func(options *cacheOptions) {
		options.poolSize = poolSize
	}
}

func MinIdleConns(minIdleConns int) Option {
	return func(options *cacheOptions) {
		options.minIdleConns = minIdleConns
	}
}

func MaxConnAge(maxConnAge time.Duration) Option {
	return func(options *cacheOptions) {
		options.maxConnAge = maxConnAge
	}
}

func MinRetryBackoff(minRetryBackoff time.Duration) Option {
	return func(options *cacheOptions) {
		options.minRetryBackoff = minRetryBackoff
	}
}

func MaxRetryBackoff(maxRetryBackoff time.Duration) Option {
	return func(options *cacheOptions) {
		options.maxRetryBackoff = maxRetryBackoff
	}
}

func DialTimeout(dialTimeout time.Duration) Option {
	return func(options *cacheOptions) {
		options.dialTimeout = dialTimeout
	}
}

func ReadTimeout(readTimeout time.Duration) Option {
	return func(options *cacheOptions) {
		options.readTimeout = readTimeout
	}
}

func WriteTimeout(writeTimeout time.Duration) Option {
	return func(options *cacheOptions) {
		options.writeTimeout = writeTimeout
	}
}

func WithCodec(codec codec.Codec) Option {
	return func(options *cacheOptions) {
		options.codec = codec
	}
}

func PoolTimeout(poolTimeout time.Duration) Option {
	return func(options *cacheOptions) {
		options.poolTimeout = poolTimeout
	}
}

func IdleTimeout(idleTimeout time.Duration) Option {
	return func(options *cacheOptions) {
		options.idleTimeout = idleTimeout
	}
}

func IdleCheckFrequency(idleCheckFrequency time.Duration) Option {
	return func(options *cacheOptions) {
		options.idleCheckFrequency = idleCheckFrequency
	}
}
