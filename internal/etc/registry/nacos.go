package registry

const Nacos = "nacos"

const NacosTemplate = `
[registry.nacos]
	# 服务器地址 [scheme://]ip:port[/nacos]。默认为["http://127.0.0.1:8848/nacos"]
	urls = ["http://127.0.0.1:8848/nacos"]
	# 集群名称。默认为DEFAULT
	clusterName = "DEFAULT"
	# 群组名称。默认为DEFAULT_GROUP
	groupName = "DEFAULT_GROUP"
	# 请求Nacos服务端超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为3秒
	timeout = "3s"
	# ACM的命名空间Id。默认为空
	namespaceId = ""
	# 当使用ACM时，需要该配置，默认为空。详见：https://help.aliyun.com/document_detail/130146.html
	endpoint = ""
	# ACM&KMS的regionId，用于配置中心的鉴权。默认为空
	regionId = ""
	# ACM&KMS的AccessKey，用于配置中心的鉴权。默认为空
	accessKey = ""
	# ACM&KMS的SecretKey，用于配置中心的鉴权。默认为空
	secretKey = ""
	# 是否开启kms，同时DataId必须以"cipher-"作为前缀才会启动加解密逻辑。kms可以参考文档：https://help.aliyun.com/product/28933.html
	openKMS = false
	# 缓存service信息的目录。默认为./run/nacos/naming/cache
	cacheDir = "./run/nacos/naming/cache"
	# Nacos服务端的API鉴权Username。默认为空
	username = ""
	# Nacos服务端的API鉴权Password。默认为空
	password = ""
	# 日志存储路径。默认为./run/nacos/naming/log
	logDir = "./run/nacos/naming/log"
	# 日志输出级别，可选：debug、info、warn、error。默认为info
	logLevel = "info"
`
