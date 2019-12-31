package main

import (
	"fmt"
	"gf-app/app/utils"
	_ "gf-app/boot"
	"gf-app/config"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	config.Init()
	urlDetail := utils.Parse("https://image.artyears.cn/image/2019-10/abc.png")
	fmt.Println(urlDetail.ToString())
	g.Server().Run()
}
