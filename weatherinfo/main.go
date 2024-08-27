package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

// const (
// 	api_key = ""
// )

type Settings struct {
	API_KEY string `yaml:"apikey"`
}

// TODO: Понять как происходит загрузка конфига и какие для этого лучшие практики

func loadConfig() (Settings, error) {
	var settings Settings
	file, err := os.Open("config.yaml")
	if err != nil {
		return settings, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		return settings, fmt.Errorf("error decoding config file: %v", err)
	}

	return settings, nil
}

type WeatherStruct struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

func CurrentWeather(city string, api_key string) {
	// Важная штука убирающая пробелы в Path-параметрах, иначе будет ошибка зачастую
	city = url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, api_key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error(), "Error in response")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		io.Copy(io.Discard, resp.Body)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error(), "Error in reading data")
		return
	}
	var weather WeatherStruct
	// fmt.Println(string(data))
	err = json.Unmarshal(data, &weather)
	if err != nil {
		fmt.Println(err.Error(), "Error in struct insert")
		return
	}
	if len(weather.Weather) > 0 {
		fmt.Printf("Today's temperature in %s is %.2f°C, weather condition is %s\n", city, weather.Main.Temp, weather.Weather[0].Main)
	} else {
		fmt.Printf("Today's temperature in %s is %.2f°C, but no weather condition data available\n", city, weather.Main.Temp)
	}
}

func main() {
	settings, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	CurrentWeather("Krasnoyarsk", settings.API_KEY)
}
