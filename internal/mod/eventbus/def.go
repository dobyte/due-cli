package eventbus

import (
	_ "embed"
)

const (
	Nats    = "nats"
	Redis   = "redis"
	Kafka   = "kafka"
	Process = "process"
)

const (
	RedisPackage = "github.com/dobyte/due/eventbus/redis"
	NatsPackage  = "github.com/dobyte/due/eventbus/nats"
	KafkaPackage = "github.com/dobyte/due/eventbus/kafka"
)

var (
	//go:embed nats.toml
	NatsTemplate string
	//go:embed redis.toml
	RedisTemplate string
	//go:embed kafka.toml
	KafkaTemplate string
)
