package transporter

const GRPC = "grpc"

const GRPCServerTemplate = `
[transport.grpc.server]
	# 服务器监听地址
	addr = ":0"
	# 秘钥文件
	keyFile = ""
	# 证书文件
	certFile = ""
`

const GRPCClientTemplate = `
[transport.grpc.client]
	# 证书文件
	certFile = ""
	# 证书域名
	serverName = ""
	# 连接池大小，默认为10
	poolSize = 10
`
