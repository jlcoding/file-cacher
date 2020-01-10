package api

import (
	"file-cacher/app/utils"
	"file-cacher/config"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Proxy(req *ghttp.Request) {
	host := GetHost(req.GetHost())
	path := req.URL.String()
	suffix := utils.GetSuffix(req.URL.String())

	bytes := utils.Download(fmt.Sprintf("http://%s%s", host, path), "")
	if nil == bytes {
		req.Response.Status = 404
	}
	// set heads
	header := req.Response.Writer.Header()
	header.Set("Content-Type", "image/"+suffix)

	_, err := req.Response.Writer.Write(bytes)
	if nil != err {
		glog.Error("response write fail", err)
	}
}

func GetHost(host string) string {
	if "" == host {
		return ""
	}
	for _, value := range config.Config.ReHost {
		if host == value.Origin {
			return value.Target
		}
	}
	glog.Error("missing rehost target host:", host)
	return ""
}
