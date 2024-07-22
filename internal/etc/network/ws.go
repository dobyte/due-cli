package network

const WS = "ws"

const WSServerTemplate = `
[network.ws.server]
	# 服务器监听地址
	addr = ":3553"
	# 客户端连接路径
	path = "/"
	# 服务器最大连接数
	maxConnNum = 5000
	# 秘钥文件
	keyFile = ""
	# 证书文件
	certFile = ""
	# 跨域检测，空数组时不允许任何连接升级成websocket，未设置此参数时允许所有的链接升级成websocket
	origins = ["*"]
	# 握手超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	handshakeTimeout = "10s"
	# 心跳检测间隔时间。设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
	# 心跳机制，默认为resp响应式心跳。可选：resp 响应式心跳 | tick 定时主推心跳
	heartbeatMechanism = "resp"
`

const WSClientTemplate = `
[network.ws.client]
	# 拨号地址
	url = "ws://127.0.0.1:3553"
	# 握手超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	handshakeTimeout = "10s"
	# 心跳间隔时间。设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
`
