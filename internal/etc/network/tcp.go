package network

const TCP = "tcp"

const TCPServerTemplate = `
[network.tcp.server]
	# 服务器监听地址
	addr = ":3553"
	# 服务器最大连接数
	maxConnNum = 5000
	# 心跳检测间隔时间。设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
	# 心跳机制，默认resp
	heartbeatMechanism = "resp"
`

const TCPClientTemplate = `
[network.tcp.client]
	# 拨号地址
	addr = "127.0.0.1:3553"
	# 拨号超时时间，默认5s
	timeout = "5s"
	# 心跳间隔时间。设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
`
