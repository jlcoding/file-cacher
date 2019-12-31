package router

import (
    "gf-app/app/api"
    "github.com/gogf/gf/frame/g"
)

// 统一路由注册.
func init() {

    g.Server().BindHandler("/download", api.GetImg)
}
