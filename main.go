package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
)

func main() {
	core.InitLogrus()
	flags.Parse()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()

	flags.Run()

	routers.Run()
}
