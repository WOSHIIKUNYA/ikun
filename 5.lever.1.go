package main

import (
	"github.com/gin-gonic/gin"
)

type user struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

func main() {
	r := gin.Default()
	r.POST("/logn", func(c *gin.Context) {
		f, err := c.Cookie("gin_cookie")
		if err != nil {
			f = "未登录"
			c.SetCookie("gin_cookie", "已登录", 36500, "/", "localhost", false, true)
		}
		c.JSON(200, f)
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
