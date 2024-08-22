package main

import "fmt"

func main() {
	n := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			n <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			n <- i
		}
		done <- true
	}()
	go func() {
		<-done
		<-done // тут два дона чтобы дождаться два канала которые выполняют прошлые две функции, и регают, важно закрыть канал, чтоб не было ошибки дедлока
		close(n)
	}()
	for a := range n {
		fmt.Println(a)
	}
}
