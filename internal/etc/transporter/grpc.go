package transporter

const GRPC = "grpc"

const GRPCServerTemplate = `
[transport.grpc.server]
	# 服务器监听地址。空或:0时系统将会随机端口号
	addr = ":0"
	# 是否将内部通信地址暴露到公网。默认为false
	expose = false
	# 私钥文件
	keyFile = ""
	# 证书文件
	certFile = ""
`

const GRPCClientTemplate = `
[transport.grpc.client]
	# CA证书文件
	caFile = ""
	# 证书域名
	serverName = ""
`
