package main

import (
	"fmt"
)

// variadict func example
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Hello World")
	sum(10, 20)
	sum(10, 20, 30)
	sum(10, 20, 30, 40)
	nums := []int{1, 3, 4, 5}
	sum(nums...)
	sum()
}
