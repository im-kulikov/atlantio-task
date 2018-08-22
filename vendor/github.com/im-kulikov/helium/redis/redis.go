package redis

import (
	"github.com/go-redis/redis"
	"github.com/im-kulikov/helium/module"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	// Config alias
	Config = redis.Options

	// Client alias
	Client = redis.Client
)

var (
	// Module is default Redis client
	Module = module.Module{
		{Constructor: NewDefaultConfig},
		{Constructor: NewConnection},
	}
	// ErrEmptyConfig when given empty options
	ErrEmptyConfig = errors.New("redis empty config")
)

func NewDefaultConfig(v *viper.Viper) *Config {
	return &Config{
		Addr:               v.GetString("redis.address"),
		Password:           v.GetString("redis.password"),
		DB:                 v.GetInt("redis.db"),
		MaxRetries:         v.GetInt("redis.max_retries"),
		MinRetryBackoff:    v.GetDuration("redis.min_retry_backoff"),
		MaxRetryBackoff:    v.GetDuration("redis.max_retry_backoff"),
		DialTimeout:        v.GetDuration("redis.dial_timeout"),
		ReadTimeout:        v.GetDuration("redis.read_timeout"),
		WriteTimeout:       v.GetDuration("redis.write_timeout"),
		PoolSize:           v.GetInt("redis.pool_size"),
		PoolTimeout:        v.GetDuration("redis.pool_timeout"),
		IdleTimeout:        v.GetDuration("redis.idle_timeout"),
		IdleCheckFrequency: v.GetDuration("redis.idle_check_frequency"),
	}
}

// New redis client
func NewConnection(opts *Config) (cache *Client, err error) {
	if opts == nil {
		return nil, ErrEmptyConfig
	}

	cache = redis.NewClient(opts)

	if _, err = cache.Ping().Result(); err != nil {
		return nil, err
	}

	return cache, nil
}
