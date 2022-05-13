package middleware

import (
	"github.com/gin-gonic/gin"
	"myproject/Model"
	"myproject/common"
	"net/http"
	"strings"
)

//生成中间件
func AuthMiddleware() gin.HandlerFunc {
	//有了token就不需要再重新输入密码了
	//我们通过验证token中用户ID来比对用户是否存在，如果存在则上传到上下文c.set
	return func(c *gin.Context) {
		//获取authorization header
		tokenstr := c.GetHeader("Authorization")

		//检验token 格式， 穿回来的格式一般都是前缀有个bearer
		if tokenstr == "" || strings.HasSuffix(tokenstr, "Bearer") { // 如果token为空或者不是以bear开头
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenstr = tokenstr[7:]
		token, claims, err := common.ParseToken(tokenstr)
		//如果解析失败或token无效
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//验证通过后claim中的userId
		userId := claims.Userid
		DB := common.GetDB()
		var user Model.User
		DB.Where("id=?", userId).First(&user)
		//用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//如果存在，将user信息存入上下文
		c.Set("user", user)
		c.Next()

	}

}
