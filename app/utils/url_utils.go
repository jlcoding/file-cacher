package utils

import (
	"fmt"
	"github.com/gogf/gf/text/gregex"
	"strings"
)

type UrlDetail struct {
	FileName string
	Suffix   string
	Host     string
	Protocol string
	Path     string
}

func (this *UrlDetail) ToString() string {
	return this.Protocol + "://" + this.Host + "/" + this.Path + "." + this.Suffix
}

/**
获取文件名
path: 文件路径
*/
func GetFileName(url string) string {
	index := strings.LastIndex(url, "/")
	return string([]rune(url)[index+1 : len(url)])
}

/**
获取文件后缀名
path: 文件路径
*/
func GetSuffix(url string) string {
	index := strings.LastIndex(url, ".")
	return string([]rune(url)[index+1 : len(url)])
}

func Parse(url string) UrlDetail {
	result := new(UrlDetail)
	match, _ := gregex.MatchString(`(http|https)://(.+?)/(.+?)\.(.+)`, url)
	fileName := GetFileName(url)

	result.Protocol = match[1]
	result.Host = match[2]
	result.Path = match[3]
	result.Suffix = match[4]
	result.FileName = fileName

	fmt.Printf("protocol:%s, host:%s fileName:%s, suffix:%s", result.Protocol, result.Host, result.FileName, result.Suffix)

	return *result
}
