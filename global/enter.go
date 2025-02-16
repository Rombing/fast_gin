package global

import (
	"fast_gin/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const Version = "1.0.0"

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
