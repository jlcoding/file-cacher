package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	ReHost []ReHost `toml:"rehost"`
}

type ReHost struct {
	Origin string `toml:"origin"`
	Target string `toml:"target"`
}

func Init()  {
	var config GlobalConfig
	const PATH = "config/config.toml"
	_, err := toml.DecodeFile(PATH, &config)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.ReHost[0])
}



