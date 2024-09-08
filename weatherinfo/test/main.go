package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/c1kvoy/repos")
	if err != nil {
		fmt.Println(err.Error(), "in response")
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error(), "in decode")
		return
	}

	// Создаем переменную для хранения декодированных данных
	var repos []interface{}

	// Декодируем JSON в переменную
	if err := json.Unmarshal(data, &repos); err != nil {
		fmt.Println(err.Error(), "in unmarshal")
		return
	}

	// Используем MarshalIndent для красивого вывода
	prettyJSON, err := json.MarshalIndent(repos, "", "  ")
	if err != nil {
		fmt.Println(err.Error(), "in marshal indent")
		return
	}

	// Выводим отформатированный JSON
	fmt.Println(string(prettyJSON))
}
