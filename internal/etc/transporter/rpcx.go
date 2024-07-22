package transporter

const RPCX = "rpcx"

const RPCXServerTemplate = `
[transport.rpcx.server]
	# 服务器监听地址，空或:0时系统将会随机端口号
	addr = ":0"
	# 秘钥文件
	keyFile = ""
	# 证书文件
	certFile = ""
`

const RPCXClientTemplate = `
[transport.rpcx.client]
	# 证书文件
	certFile = ""
	# 证书域名
	serverName = ""
	# 连接池大小，默认为10
	poolSize = 10
`
