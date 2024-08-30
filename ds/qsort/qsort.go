package main

import "fmt"
// реализация qsort на голанге
func Qsort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		} else {
			return []int{arr[0], arr[1]}
		}
	}
	pivot := arr[0]
	var below []int
	var above []int
	for _, num := range arr[1:] {
		if num <= pivot {
			below = append(below, num)
		} else {
			above = append(above, num)
		}
	}
	return append(append(Qsort(below), pivot), Qsort(above)...)
}
func main() {
	nonsorted := []int{23123, 2, 3, 1, 5}
	fmt.Println(Qsort(nonsorted))
}

