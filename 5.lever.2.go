package main

import (
	"github.com/gin-gonic/gin"
)

var a [100]person
var m int = 0
var t int = 0
var h bool

type person struct {
	Name     string `json:"Name"`
	Possword string `json:"Possword"`
}

//定义了一个结构体
func main() {
	r := gin.Default() //打开引擎

	r.POST("/pin", func(c *gin.Context) {
		c.ShouldBind(&a[t])
		c.JSON(200, a[t])
		t++
	})

	//启动注册服务路径为/pin
	r.POST("/c", func(c *gin.Context) {
		var x person
		c.ShouldBind(&x)
		h = false
		for f := 0; f <= t; f++ {
			if x.Possword == a[f].Possword || x.Name == a[f].Name {
				f = t + 1
				h = true
			}
		}

		//开启登录服务路径为/c
		if h == true {
			c.JSON(200, x)
		} else {
			c.String(200, "密码错误")
		}
	})

	//密码服务
	r.Run()
}
