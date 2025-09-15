package cache

import (
	_ "embed"
)

const (
	Redis    = "redis"
	Memcache = "memcache"
)

var (
	//go:embed redis.toml
	RedisTemplate string
	//go:embed memcache.toml
	MemcacheTemplate string
)
