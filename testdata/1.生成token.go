package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/utils/jwts"
	"fmt"
)

func main() {
	core.InitLogrus()
	flags.Parse()
	global.Config = core.ReadConfig()
	token, err := jwts.SetToken(jwts.Claims{
		UserID: 1,
		RoleID: 1,
	})
	fmt.Println(token, err)
}
