package api

import (
	"file-cacher/app/utils"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type Remote struct {
}

func (*Remote) Get(req *ghttp.Request) {
	url := req.GetString("url")
	sha1 := req.GetString("sha1")
	bytes := utils.Download(url, sha1)
	if nil == bytes {
		req.Response.Status = 404
	}

	// set heads
	header := req.Response.Writer.Header()
	fileType := utils.GetFileType(bytes)
	if "" == fileType {
		fileType = "application/octet-stream"
	}
	glog.Info("fileType => " + fileType)
	header.Set("Content-Type", fileType)

	_, err := req.Response.Writer.Write(bytes)
	if nil != err {
		glog.Error("response write fail", err)
	}
}
