package boot

import (
	"file-cacher/config"
	"github.com/gogf/gf/frame/g"
)

func init() {
	g.Server().SetAddr(config.Config.Server.Address)
	g.Server().SetPort(config.Config.Server.Port)
}
