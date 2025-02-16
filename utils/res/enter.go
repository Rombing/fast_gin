package res

import (
	"fast_gin/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Ok(data, "成功", c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	Ok(map[string]any{
		"list":  list,
		"count": count,
	}, "成功", c)
}
func Fail(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(7, msg, c)
}

func FailWithError(err error, c *gin.Context) {
	msg := validate.ValidateError(err)
	Fail(7, msg, c)
}
