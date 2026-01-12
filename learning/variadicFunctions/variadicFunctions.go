package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, "")
	total:= 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main() {
	sum(1, 2, 3, 4, 5)
	sum(1, 2)

	nums := []int{1, 4, 7, 8, 4, 1, 1, 1, 1}

	sum(nums...)


}