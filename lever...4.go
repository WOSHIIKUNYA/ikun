package main

import (
	"fmt"
)

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			<-over
		}()
		over <- true
	}
	fmt.Println("over!!!")
}

/*func main(){
over:=make(chan bool)
for i:=0;i<10;i++{
go func(){
fmt.Println(i)
}（）//这是携程已经开始进行一共先后开了9个携程
if i==9{
over<-true
}//当i=9时主程序堵塞，此时已经开了9个携程
<-over//无效操作主携程以被堵塞接受也无法进行
fmt.Print(OVER!!!)
}















*/
