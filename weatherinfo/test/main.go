package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/")
	if err!=nil {
		fmt.Println(err.Error(), "in response")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error(), "in decode")
	}
	fmt.Println(string(data))
}