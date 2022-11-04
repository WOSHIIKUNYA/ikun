package main

import (
	"fmt"
	"sync"
)

var x int
var b = make(chan int, 2)
var mu sync.WaitGroup

func add() {
	for i := 0; i < 50000; i++ {
		b <- i
		x = x + 1
		<-b
	}
	mu.Done()
}

func main() {
	mu.Add(2)
	go add()
	go add()

	mu.Wait()
	fmt.Println(x)
}
