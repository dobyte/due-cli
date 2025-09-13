package base

const Template = `
# 进程号
pid = "./run/due.pid"
# 开发模式。支持模式：debug、test、release（设置优先级：配置文件 < 环境变量 < 运行参数 < mode.SetMode()）
mode = "debug"
# 统一时区设置。项目中的时间获取请使用xtime.Now()
timezone = "Local"
# 容器关闭最大等待时间。支持单位：纳秒（ns）、微秒（us | µs）、毫秒（ms）、秒（s）、分（m）、小时（h）、天（d）。默认为0
shutdownMaxWaitTime = "0s"
`
