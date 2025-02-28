package template

const EtcOutput = "etc/etc.toml"

const EtcTemplate = `
# 进程号
pid = "./run/${VarName}.pid"
# 统一时区设置。项目中的时间获取请使用xtime.Now()
timezone = "Local"

# 任务池
[task]
    # 任务池大小(goroutine)
    size = 100000
    # 是否非阻塞
    nonblocking = true
    # 是否禁用清除。
    disablePurge = true

[cluster]
    # 集群网格配置
    [cluster.mesh]
        # 实例ID，网关集群中唯一。不填写默认自动生成唯一的实例ID
        id = ""
        # 实例名称
        name = "${VarName}"
        # 编解码器。可选：json | proto。默认json
        codec = "json"
        # RPC调用超时时间，支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为3s
        timeout = "3s"

# 框架默认打包器统一采用以下的打包格式，自定义打包器可自行定义打包规则
# 心跳包
# ------------------------------------------------------------------------------
# | size(4 byte) = (1 byte + 8 byte) | header(1 byte) | heartbeat time(8 byte) |
# ------------------------------------------------------------------------------
#
#  0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
# +---------------------------------------------------------------+-+-------------+---------------------------------------------------------------+
# |                              size                             |h|   extcode   |                      heartbeat time (ns)                      |
# +---------------------------------------------------------------+-+-------------+---------------------------------------------------------------+

# 数据包
# -----------------------------------------------------------------------------------------------------------------------
# | size(4 byte) = (1 byte + n byte + m byte + x byte) | header(1 byte) | route(n byte) | seq(m byte) | message(x byte) |
# -----------------------------------------------------------------------------------------------------------------------
#
#  0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
# +---------------------------------------------------------------+-+-------------+-------------------------------+-------------------------------+
# |                              size                             |h|   extcode   |             route             |              seq              |
# +---------------------------------------------------------------+-+-------------+-------------------------------+-------------------------------+
# |                                                                message data ...                                                               |
# +-----------------------------------------------------------------------------------------------------------------------------------------------+

[packet]
    # 字节序，默认为big。可选：little | big
    byteOrder = "big"
    # 路由字节数，默认为2字节
    routeBytes = 2
    # 序列号字节数，默认为2字节
    seqBytes = 2
    # 消息字节数，默认为5000字节
    bufferBytes = 5000

[log]
    # 日志输出文件
    file = "./log/due.log"
    # 日志输出级别，可选：debug | info | warn | error | fatal | panic
    level = "info"
    # 日志输出格式，可选：text | json
    format = "text"
    # 是否输出到终端
    stdout = true
    # 时间格式，标准库时间格式
    timeFormat = "2006/01/02 15:04:05.000000"
    # 堆栈的最低输出级别，可选：debug | info | warn | error | fatal | panic
    stackLevel = "error"
    # 文件最大留存时间，d:天、h:时、m:分、s:秒
    fileMaxAge = "7d"
    # 文件最大尺寸限制，单位（MB）
    fileMaxSize = 100
    # 文件切割方式
    fileCutRule = "day"
    # 是否启用调用文件全路径
    callerFullPath = true
    # 是否启用分级存储
    classifiedStorage = true

${VarLocatorEtcBlock}



`
