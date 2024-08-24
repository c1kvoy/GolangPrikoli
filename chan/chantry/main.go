package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan string, 2)
	go ParsSimul(ch1)
	time.Sleep(3 * time.Second)
	select {
	case <-ch1:
		fmt.Println("Success")
	default:
		fmt.Println("Nothing")
	}
}

func ParsSimul(ch chan string) {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
	a := rand.Intn(2)
	if a == 0 {
		ch <- "Parsing Success"
		ch <- "Info"
		return
	} else {
		return
	}
}
