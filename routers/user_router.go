package routers

import (
	"fast_gin/api"
	"fast_gin/api/user_api"
	"fast_gin/middleware"
	"fast_gin/models"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	app := api.App.UserApi
	g.POST("users/login",
		middleware.LimitMiddleware(1),
		middleware.BindJsonMiddleware[user_api.LoginRequest],
		app.LoginView)
	g.GET("users",
		middleware.LimitMiddleware(10),
		middleware.AuthMiddleware,
		middleware.BindQueryMiddleware[models.PageInfo],
		app.UserListView)
	g.POST("users/logout",
		middleware.LimitMiddleware(1),
		middleware.AuthMiddleware,
		app.LogoutView)
}
