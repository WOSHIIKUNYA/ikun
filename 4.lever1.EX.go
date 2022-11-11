package main

import (
	"fmt"
	"strings"
)

func bb() {
	fmt.Println("按1修改技能1")
	fmt.Println("按2修改技能2")
	fmt.Println("按3修改技能3")
	fmt.Println("按4输出模板，技能名字用skill替代")
	fmt.Println("按5释放技能x")
	fmt.Println("按0退出")
}

// 菜单功能
func MingGanCi(b string) bool {
	c := strings.Contains(b, "你妈")
	c = strings.Contains(b, "傻比")
	c = strings.Contains(b, "强奸")
	if c == true {
		return c
	} else {
		return false
	}
}

// 查找功能

func main() {
	var b int
	c := make([]string, 3)
	//c是一个切片包含了3个技能名
	var d string
	c[0] = "技能1"
	c[1] = "技能2"
	c[2] = "技能3"

	//初始化三个技能名
	for {
		bb()
		//这个是菜单的函数
		fmt.Scanln(&b)
		switch b {
		case 1: //写技能1
			fmt.Println("请输入技能名")
			fmt.Scanln(&c[0])
			x := MingGanCi(c[0])
			if x == false {
				fmt.Println("修改技能成功")
			} else {
				fmt.Println("修改技能失败，包含敏感词")
				c[0] = "技能1"
			}
			//选择你的功能
		case 2:
			fmt.Println("请输入技能名")
			fmt.Scanln(&c[1])
			x := MingGanCi(c[1])
			if x == false {
				fmt.Println("修改技能成功")
			} else {
				fmt.Println("修改技能失败，包含敏感词")
				c[1] = "技能2"
			}
			//写技能2
		case 3:
			fmt.Println("请输入技能名")
			fmt.Scanln(&c[2])
			x := MingGanCi(c[2])
			if x == false {
				fmt.Println("修改技能成功")
			} else {
				fmt.Println("修改技能失败，包含敏感词")
				c[2] = "技能3"
			}
			//写技能3
		case 4:

			d = "尝尝我的厉害"
			//初始化模板
			fmt.Println("请输入模板")
			fmt.Scanln(&d)
			x := MingGanCi(d)
			if x == false {
				fmt.Println("修改模板成功")
			} else {
				fmt.Println("修改模板失败，包含敏感词")
			}
			//写模板
		case 5:
			var a int
			for g := 0; g < 3; g++ {
				fmt.Println(g+1, ":", c[g])
			}
			fmt.Println("请输入你要释放的技能")
			fmt.Scanln(&a)
			ReleaseSkill1(c[a-1], d, func(skillName string, moban string) {
				f := strings.Replace(moban, "skill", skillName, -1) //替换
				fmt.Println(f)
			})

		case 0:
			return
			fmt.Println("谢谢使用")
		}
	}
}
func ReleaseSkill1(a string, b string, c func(string, string)) {
	c(a, b)
}

//把a技能名，b模板进去
