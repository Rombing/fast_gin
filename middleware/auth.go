package middleware

import (
	"fast_gin/service/redis_ser"
	"fast_gin/utils/jwts"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", c)
		c.Abort()
		return
	}
	if redis_ser.HasLogout(token) {
		res.FailWithMsg("用户已注销", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", c)
		c.Abort()
		return
	}
	if redis_ser.HasLogout(token) {
		res.FailWithMsg("用户已注销", c)
		c.Abort()
		return
	}
	if claims.RoleID != 1 {
		res.FailWithMsg("角色认证失败", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}

func GetAuth(c *gin.Context) (cl *jwts.Claims) {
	cl = new(jwts.Claims)
	_claims, ok := c.Get("claims")
	if !ok {
		return
	}
	cl, ok = _claims.(*jwts.Claims)
	return
}
