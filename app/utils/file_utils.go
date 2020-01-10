package utils

import (
	_const "file-cacher/app/common"
	"file-cacher/app/db"
	"file-cacher/config"
	"fmt"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"os"
	"strconv"
	"strings"
)

func WriteToFile(path string, bytes []byte, suffix string) int {
	pathParts := strings.Split(path, "/")
	pathParts = pathParts[1:]
	var file *os.File
	var err error
	var size int
	var tempPath string

	for index, value := range pathParts {
		if "" != suffix {
			suffix := GetSuffix(value)
			value = strings.Replace(value, "."+suffix, ".gzip", 1)
		}

		tempPath += "/" + value

		if !gfile.Exists(tempPath) && !gfile.IsDir(tempPath) {
			if index >= (cap(pathParts) - 1) {
				file, err = gfile.Create(tempPath)
				if nil != err {
					glog.Error("Create File Fail", err)
				}
				break
			} else {
				err = gfile.Mkdir(tempPath)
				if nil != err {
					glog.Error("Create File Fail", err)
				}
			}
		}
	}

	size, err = file.Write(bytes)
	if nil != err {
		glog.Error("write file err", err)
		glog.Info("write file result" + strconv.Itoa(size))
	}
	defer file.Close()
	return size
}

func Download(url string, sha1 string) []byte {
	var bytes []byte
	if "" != sha1 {
		path := db.Get(_const.BuildKey([]string{_const.SHA1_PATH, sha1}))
		if "" != path {
			if gfile.Exists(path) {
				bytes = gfile.GetBytes(path)
				return bytes
			}
		}
	}

	if "" == url {
		glog.Error(" url is null, can't get file and no cache at disk")
		return nil
	}

	// 发送GET方法
	client := ghttp.NewClient()
	bytes = client.GetBytes(url)
	if nil == bytes {
		return nil
	}
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	filePath := fmt.Sprintf(config.Config.DataPath+"/%s", url)

	if gfile.Exists(filePath) {
		bytes = gfile.GetBytes(filePath)
	} else {
		_ = WriteToFile(filePath, bytes, "")
	}
	updateSha1(bytes, filePath)
	return bytes
}

func updateSha1(bytes []byte, filePath string) {
	sha1 := gsha1.Encrypt(bytes)
	path := db.Get(_const.BuildKey([]string{_const.SHA1_PATH, sha1}))
	if "" == path {
		db.Update(_const.BuildKey([]string{_const.SHA1_PATH, sha1}), filePath)
	}
}
