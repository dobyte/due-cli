package locate

import (
	_ "embed"
)

const (
	Redis = "redis"
)

const (
	RedisPackage = "github.com/dobyte/due/locate/redis"
)

var (
	//go:embed redis.toml
	RedisTemplate string
)
