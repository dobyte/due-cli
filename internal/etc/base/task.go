package base

const TaskTemplate = `
# 任务池
[task]
    # 任务池大小(goroutine)
    size = 100000
    # 是否非阻塞
    nonblocking = true
    # 是否禁用清除。
    disablePurge = true
`
