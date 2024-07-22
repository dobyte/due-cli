package cluster

const Mesh = "mesh"

const MeshTemplate = `
[cluster.mesh]
	# 实例ID，网关集群中唯一。不填写默认自动生成唯一的实例ID
	id = ""
	# 实例名称
	name = "${VarName}"
	# 编解码器。可选：json | proto。默认json
	codec = "${VarCodec}"
	# RPC调用超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为3s
	timeout = "3s"
`
