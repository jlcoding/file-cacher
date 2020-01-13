package config

import (
	"github.com/gogf/gf/encoding/gtoml"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

type GlobalConfig struct {
	DataPath string   `toml:"data-path"`
	ReHost   []ReHost `toml:"rehost"`
	Server   Server   `toml:"server"`
	Proxy    Proxy    `toml:"proxy"`
	BadgerDB BadgerDB `toml:"badgerdb"`
}

/**
  host替换
*/
type ReHost struct {
	Origin string `toml:"origin"`
	Target string `toml:"target"`
}

type Path struct {
	Data string `toml:"data"`
}

type Server struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

type Proxy struct {
	ProxyMode bool     `toml:"proxy-mode"`
	ProxyHost []string `toml:"proxy-host"`
	Suffix    []string `toml:"suffix"`
}

type BadgerDB struct {
	Dir      string `toml:"dir"`
	ValueDir string `toml:"value-dir"`
}

var Config GlobalConfig

func init() {
	glog.Info("config init...")
	bytes := gfile.GetBytes("config/config.toml")
	err := gtoml.DecodeTo(bytes, &Config)

	if err != nil {
		glog.Error(err)
		return
	}
}
