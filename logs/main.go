package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const(
	api_key = "bcf5e86142371b09b2bd3a766310da2b"
)
func main() {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", "Moscow", api_key)
	resp,err := http.Get(url)
	if err != nil{
		log.Println("Ошибка в запросе данных")
	}
	defer resp.Body.Close()
	log.Println("All good")
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(string(data))
}