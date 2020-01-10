package main

import (
	_ "file-cacher/config"
	_ "file-cacher/boot"
	_ "file-cacher/router"
	_ "file-cacher/app/db"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
