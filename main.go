package main

import (
	_ "file-cacher/app/db"
	_ "file-cacher/boot"
	_ "file-cacher/config"
	_ "file-cacher/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
