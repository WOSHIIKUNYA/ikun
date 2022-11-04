package main

import (
	"fmt"
	"os"
)

func write(a string) {
	f, err := os.Create(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("I'm not  afraid of difficulties and insist on learning programming")
}
func read(a string) {
	f, err := os.Open(a)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	b := make([]byte, 100)
	n, err1 := f.Read(b)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	println(string(b[0:n]))

}
func main() {
	path := "./plan.txt"
	write(path)
	read(path)
}
