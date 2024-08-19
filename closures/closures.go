package main

import "fmt"

func closure() func() string {
	var rap string = ""
	fmt.Println(rap)
	return func() string {
		rap += "a"
		return rap
	}
}

func main() {
	test := closure()
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	
}
