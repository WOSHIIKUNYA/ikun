package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(300 * time.Second)
	go func() {
		for {
			<-ticker.C
			fmt.Println("坤瘾犯了")
		}
	}()
	stduy := time.NewTicker(2 * time.Hour)
	go func() {
		for {
			<-stduy.C
			fmt.Println("快去学习")
		}
	}()
	fly := time.NewTicker(30 * time.Second)
	go func() {
		for {
			<-fly.C
			fmt.Println("芜湖，起飞")
		}
	}()
	time.Sleep(24 * time.Hour)
}
