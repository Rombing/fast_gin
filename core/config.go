package core

import (
	"fast_gin/config"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile("settings.yaml")
	if err != nil {
		fmt.Printf("Read settings file failed, err:%v\n", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		fmt.Printf("Parse settings file failed, err:%v\n", err)
		return
	}
	return
}
