package template

const CoreOutput = "app/logic/core.go"

const CoreTemplate = `
package logic

import (
	"${VarName}/app/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

type core struct {
	proxy *node.Proxy
}

func NewCore(proxy *node.Proxy) *core {
	return &logic{
		proxy: proxy,
	}
}

func (c *core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册中间件
		group.Middleware(middleware.Auth)
		// 注册路由
		group.AddRouteHandler(route.Greet, false, c.greetHandler)
	})

	// 断线重连
	c.proxy.AddEventHandler(cluster.Reconnect, c.reconnect)
	// 断开连接
	c.proxy.AddEventHandler(cluster.Disconnect, c.disconnect)
}

// 重新连接
func (c *core) reconnect(ctx node.Context) {
	log.Debugf("client reconnect, cid: %d uid: %d", ctx.CID(), ctx.UID())
}

// 断开连接
func (c *core) disconnect(ctx node.Context) {
	log.Debugf("client disconnect, cid: %d uid: %d", ctx.CID(), ctx.UID())
}

// 路由消息处理器
func (c *core) greetHandler(ctx node.Context) {
	req := &GreetReq{}
	res := &GreetRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response message failed: %v", err)
		}
	}()

	if err := ctx.Parse(req); err != nil {
		log.Errorf("parse request message failed: %v", err)
		return
	}

	log.Info(req.Message)

	res.Message = fmt.Sprintf("I'm server, and the current time is: %s", xtime.Now().Format(xtime.DatetimeLayout))
}
`
