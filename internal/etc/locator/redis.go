package locator

const Redis = "redis"

const RedisTemplate = `
[locate.redis]
	# 客户端连接地址
	addrs = ["127.0.0.1:6379"]
	# 数据库号
	db = 0
	# 用户名
	username = ""
	# 密码
	password = ""
	# 最大重试次数
	maxRetries = 3
	# key前缀
	prefix = "due"
`
