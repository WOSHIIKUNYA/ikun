package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// 我用的是国产的Apifox传文件不会用就用了io流来编辑
type per struct {
	Name     string `json:"Name"`
	Possword string `json:"Possword"`
}

var qq bool = false

func pp(x *gin.Context) {
	l, err3 := os.ReadFile("users.txt")
	if err3 != nil {
		fmt.Println(err3)
	}
	hh := make(map[string]string)
	err4 := json.Unmarshal(l, &hh)
	if err4 != nil {
		fmt.Println(err4)
	}
	var rr per
	x.ShouldBind(&rr)
	hh[rr.Name] = rr.Possword
	y, err5 := os.Create("users.txt")
	if err5 != nil {
		fmt.Println(err5)
	}
	aa, _ := json.Marshal(hh)
	y.WriteString(string(aa))
	defer y.Close()
	x.String(200, "已注册")
}

func main() {
	qq = false
	r := gin.Default()

	r.POST("/chick", pp)
	r.POST("/ee", func(c *gin.Context) {
		l, err3 := os.ReadFile("users.txt")
		if err3 != nil {
			fmt.Println(err3)
		}
		v := make(map[string]string)
		err4 := json.Unmarshal(l, &v)
		c.String(200, string(l))
		c.JSON(200, v)
		if err4 != nil {
			fmt.Println(err4)
		}
		var rr per
		if rr.Possword == v[rr.Name] {
			c.String(200, "已登录")
			c.JSON(200, rr)
		} else {
			c.String(200, "登录失败")
		}
	})
	r.Run()
}
