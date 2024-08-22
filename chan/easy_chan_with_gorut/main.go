package main

import (
	"fmt"
	"io"
	"net/http"
	_ "time"
)

func main() {
	chl := make(chan string, 2)
	go req(chl)
	fmt.Println(<-chl)
}

func req(chl chan string) {
	response, err := http.Get("https://youtube.com/")
	if err != nil {
		chl <- err.Error()
		return
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		chl <- fmt.Sprintf("Error: received status code %d", response.StatusCode)
		return
	}
	chl <- fmt.Sprintf("All okey, data: %s", data[:50])
}
