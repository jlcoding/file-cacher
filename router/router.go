package router

import (
	"file-cacher/app/api"
	"file-cacher/config"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// 统一路由注册.
func init() {
	glog.Info("init router...")
	remote := new(api.Remote)
	g.Server().BindObject("/{.struct}/{.method}", remote)

	if config.Config.Proxy.ProxyMode {
		g.Server().BindHandler("/*", api.Proxy)
	}
}
