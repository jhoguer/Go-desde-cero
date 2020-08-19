package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 5, 7))
}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}
