package network

const TCP = "tcp"

const TCPServerTemplate = `
[network.tcp.server]
	# 服务器监听地址
	addr = ":3553"
	# 私钥文件
	keyFile = ""
	# 证书文件
	certFile = ""
	# 服务器最大连接数
	maxConnNum = 5000
	# 心跳间隔时间；设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
	# 心跳机制，默认resp
	heartbeatMechanism = "resp"
	# 授权超时时间，（在客户端建立连接后，如果在授权超时时间内未进行绑定用户操作，则被认定为未授权连接，服务器会强制断开连接）支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为0s，不进行授权检测
	authorizeTimeout = "0s"
`

const TCPClientTemplate = `
[network.tcp.client]
	# 拨号地址
	addr = "127.0.0.1:3553"
	# 拨号超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为5s
	timeout = "5s"
	# CA证书文件
	caFile = ""
	# 证书域名
	serverName = ""
	# 心跳间隔时间；设置为0则不启用心跳检测，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为10s
	heartbeatInterval = "10s"
`
