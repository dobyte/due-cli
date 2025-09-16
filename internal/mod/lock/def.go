package lock

import (
	_ "embed"
)

const (
	Redis    = "redis"
	Memcache = "memcache"
)

const (
	RedisPackage    = "github.com/dobyte/due/lock/redis"
	MemcachePackage = "github.com/dobyte/due/lock/memcache"
)

var (
	//go:embed redis.toml
	RedisTemplate string
	//go:embed memcache.toml
	MemcacheTemplate string
)
