package cluster

const Node = "node"

const NodeTemplate = `
# 集群节点配置
[cluster.node]
	# 实例ID，集群中唯一。不填写默认自动生成唯一的实例ID
	id = ""
	# 实例名称
	name = "${VarName}"
	# 内建RPC服务器监听地址。不填写默认随机监听
	addr = ":0"
	# 是否将内部通信地址暴露到公网。默认为false
	expose = false
	# 编解码器。可选：json | proto。默认为proto
	codec = "${VarCodec}"
	# RPC调用超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为3s
	timeout = "3s"
	# 节点权重，用于节点无状态路由消息的加权轮询策略，权重值必需大于0才生效。默认为1
	weight = 1
	# 实例元数据
	[cluster.node.metadata]
		# 键值对，且均为字符串类型。由于注册中心的元数据参数限制，建议将键值对的数量控制在20个以内，键的字符长度控制在127个字符内，值得字符长度控制在512个字符内。
		key = "value"
`
