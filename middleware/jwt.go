package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //无权限，token是无权限的，是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //Token无效
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{ //map[string]interface{}的别名。在 Gin 中，常用于构建 JSON 响应的数据结构
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort() //终止当前请求的处理流程，阻止后续的中间件和处理函数继续执行
			return
		}
		c.Next() //用于将控制权传递给下一个中间件或处理函数
	}
}
