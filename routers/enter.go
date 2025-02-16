package routers

import (
	"fast_gin/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	gin.SetMode(global.Config.System.Mode)

	r := gin.Default()
	r.Static("/uploads", "uploads")

	g := r.Group("api")
	UserRouter(g)
	ImageRouter(g)
	CaptchaRouter(g)

	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		logrus.Info("后端服务运行在 ", addr)
	}
	r.Run(addr)
}
