package main

import (
	"fmt"
	"time"
)

func main() {
	success := make(chan int)
	bad := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-success)
			time.Sleep(1 * time.Second)
		}
		bad <- 0
	}()
	TakeOne(success, bad)
}

func TakeOne(success, bad chan int) {
	x := 0
	for{
		select {
		case success <- x:
			x += 1
		case <-bad:
			fmt.Println("exit")
			return
		}
	}
}
	
