package main

import (
	"fmt"
	"sync"
)

var a sync.WaitGroup
var c sync.Mutex

func scren() {
	for i := 1; i < 101; {
		fmt.Println(i)
		c.Lock()
		i++
		c.Unlock()
	}
	a.Done()
}
func main() {
	a.Add(2)
	go scren()
	go scren()
	a.Wait()
}
