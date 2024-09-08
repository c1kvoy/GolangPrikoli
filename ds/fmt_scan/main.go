package main

import "fmt"

func After(nums []int) bool {
	hm := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, exist := hm[num]; exist {
			return true
		}
		hm[num] = true
	}
	return false
}
func main() {
	// l := []int{1, 2, 3, 5}
	var size int
	fmt.Scan(&size)
	nums := make([]int, size)
	for i:=0; i<size; i++{
		fmt.Scan(&nums[i])
	}
	fmt.Println(After(nums))
}
