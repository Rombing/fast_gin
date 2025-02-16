package core

import (
	"fast_gin/config"
	"fast_gin/flags"
	"fast_gin/global"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("Read settings file failed, err:%v\n", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		logrus.Fatalf("Parse settings file failed, err:%v\n", err)
		return
	}
	logrus.Infof("%s Read settings file success", flags.Options.File)
	return
}

func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Errorf("Parse settings file failed, err:%v\n", err)
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("Write settings file failed, err:%v\n", err)
		return
	}
	logrus.Infof("Write settings file success")
	return
}
