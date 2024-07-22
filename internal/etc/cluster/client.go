package cluster

const Client = "client"

const ClientTemplate = `
[cluster.client]
	# 实例ID，网关集群中唯一。不填写默认自动生成唯一的实例ID
	id = ""
	# 实例名称
	name = "${VarName}"
	# 编解码器。可选：json | proto
	codec = "${VarCodec}"
`
