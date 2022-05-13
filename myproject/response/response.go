package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回一个相应
func Response(c *gin.Context, httpstatus int, code int, data gin.H, msg string) {
	c.JSON(httpstatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
