package core

import (
	"fast_gin/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"time"
)

func InitGorm() (database *gorm.DB) {
	cfg := global.Config.DB
	var dialector = cfg.DSN()
	if dialector == nil {
		return
	}

	database, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成实体外键
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
		return
	}
	// 配置连接池
	sqlDB, err := database.DB()
	if err != nil {
		logrus.Fatalf("获取数据库连接失败 %s", err)
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
		return
	}
	// 设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功")
	return
}
