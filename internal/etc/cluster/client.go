package cluster

const Client = "client"

const ClientTemplate = `
# 集群客户端配置，常用于调试使用
[cluster.client]
	# 实例ID，集群中唯一。不填写默认自动生成唯一的实例ID
	id = ""
	# 实例名称
	name = "${VarName}"
	# 编解码器。可选：json | proto。默认为proto
	codec = "${VarCodec}"
`
