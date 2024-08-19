package main 

import "fmt"

func sum(arr []int) int{
	sum:=0 
	for _, i:= range arr{
		sum += i
	}
	return sum
}

func main(){
	num := 1233;
	if num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
