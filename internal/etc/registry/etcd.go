package registry

const Etcd = "etcd"

const EtcdTemplate = `
[registry.etcd]
	# 客户端连接地址，默认为["127.0.0.1:2379"]
	addrs = ["127.0.0.1:2379"]
	# 客户端拨号超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为5s
	dialTimeout = "5s"
	# 命名空间，默认为services
	namespace = "services"
	# 超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为3s
	timeout = "3s"
	# 心跳重试次数，默认为3
	retryTimes = 3
	# 心跳重试间隔，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	retryInterval = "10s"
`
