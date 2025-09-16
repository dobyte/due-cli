package cache

import (
	_ "embed"
)

const (
	Redis    = "redis"
	Memcache = "memcache"
)

const (
	RedisPackage    = "github.com/dobyte/due/cache/redis"
	MemcachePackage = "github.com/dobyte/due/cache/memcache"
)

var (
	//go:embed redis.toml
	RedisTemplate string
	//go:embed memcache.toml
	MemcacheTemplate string
)
