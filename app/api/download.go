package api

import (
	"gf-app/app/utils"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"os"
	"strconv"
)

func GetImg(req *ghttp.Request) []byte{
	url := req.GetString("imgUrl")
	// 发送GET方法
	client := ghttp.NewClient()

	bytes := client.GetBytes(url)

	fileDetail := utils.Parse(url)

	localPath := "/Users/jiaqi/gofile/" + fileDetail.FileName


	var file *os.File
	if gfile.Exists(localPath) {
		gfile.Info("file already exist")
		var err error
		file, err = gfile.Open(localPath)
		if nil != err {
			glog.Error("open file err", err)
		}

		bytes = []byte{}
		file.Read(bytes)
	}else {
		var err error
		file, err = gfile.Create(localPath)
		if nil != err {
			glog.Error("create file err", err)

		}

		result, err := file.Write(bytes)
		if nil != err {
			glog.Error("write file err", err)
			glog.Info("write file result" + strconv.Itoa(result))
		}else {
			glog.Info("write file result" + strconv.Itoa(result))
		}
	}
	req.Response.Writer.Header().Set("Content-Type", "image/png")
	return bytes
}


