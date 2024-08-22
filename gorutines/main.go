package main

import (
	"fmt"
	"net/http"
	"sync"
	_ "time"
)

func main() {
	urls := []string{
		"https://google.com/",
		"https://youtube.com/",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			rq(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	fmt.Println("Main procces terminated")
}

func rq(url string) {
	resp, error := http.Get(url)
	if error != nil {
		fmt.Printf("Failed to get data from %s , status code: %s \n", url, error.Error())
	}
	defer resp.Body.Close()
	fmt.Printf("Data got from %s , status code: %d \n", url, 200)
}
