package registry

const Consul = "consul"

const ConsulTemplate = `
[registry.consul]
	# 客户端连接地址，默认为127.0.0.1:8500
	addr = "127.0.0.1:8500"
	# 是否启用健康检查，默认为true
	healthCheck = true
	# 健康检查时间间隔（秒），仅在启用健康检查后生效，默认为10
	healthCheckInterval = 10
	# 健康检查超时时间（秒），仅在启用健康检查后生效，默认为5
	healthCheckTimeout = 5
	# 是否启用心跳检查，默认为true
	heartbeatCheck = true
	# 心跳检查时间间隔（秒），仅在启用心跳检查后生效，默认为10
	heartbeatCheckInterval = 10
	# 健康检测失败后自动注销服务时间（秒），默认为30
	deregisterCriticalServiceAfter = 30
`
