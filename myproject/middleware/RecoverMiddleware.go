package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myproject/response"
)

func RecoveryMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(c, nil, fmt.Sprint(err))
			}
		}()
		c.Next()
	}
}
