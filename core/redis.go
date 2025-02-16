package core

import (
	"context"
	"fast_gin/global"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func InitRedis() (client *redis.Client) {
	cfg := global.Config.Redis
	if cfg.Addr == "" {
		logrus.Warnf("redis配置为空")
		return
	}
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("连接redis失败 %s", err)
		return
	}
	logrus.Infof("成功连接redis")
	return
}
