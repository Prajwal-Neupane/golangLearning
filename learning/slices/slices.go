package main

import "fmt"

func main() {


	var nums = make([]int, 2, 3) // making the initial size of the slice to 2


	nums = append(nums, 3, 4, 5 , 1)

	fmt.Println((nums))

	fmt.Println(cap(nums))

	fmt.Println(nums[2:])


}