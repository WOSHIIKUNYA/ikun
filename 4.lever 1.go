package main

import "fmt"

func main() {
	var b string
	fmt.Println("请输入模板")
	fmt.Scanln(&b)
	fmt.Println("请输入技能名")
	var a string
	fmt.Scanln(&a)
	ReleaseSkill(a, func(skillName string) {
		fmt.Println(b, skillName)
	})
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
