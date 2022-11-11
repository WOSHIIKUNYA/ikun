package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func caodan() {
	fmt.Println("1.修改信息")
	fmt.Println("2.保存创建")
}
func click(a string, b int) bool {
	x := true
	if b < 100000 {
		fmt.Println("密码过短")
		x = false
	}
	return x
}
func main() {
	//读取操作
	l, err3 := os.ReadFile("users.txt")
	if err3 != nil {
		fmt.Println(err3)
	}
	h := make(map[string]int)
	fmt.Println(string(l))
	err4 := json.Unmarshal(l, &h)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(h)
	a, err1 := os.Create("users.txt")
	var x string
	var r int
	for jj := 0; jj == 0; {
		caodan()
		var dd int
		fmt.Println("请输入选择")
		fmt.Scanln(&dd)
		switch dd {
		case 1:
			fmt.Println("请添加用户名")
			fmt.Scanln(&x)
			fmt.Println("请添加密码")
			fmt.Scanln(&r)
			gg := click(x, r)
			if gg == false {
				continue
			}
		case 2:
			fmt.Println("以完成填写")
			jj = 1
		}
	}
	h[x] = r
	m, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	//健是用户名，值是密码
	if err1 != nil {
		fmt.Println(err1)
	}
	v, err2 := a.WriteString(string(m))
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(v)
	defer a.Close()

	fmt.Println("已完成")
}
