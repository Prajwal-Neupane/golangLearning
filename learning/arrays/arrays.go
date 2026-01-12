package main

import "fmt"

func main() {


	// var intArr [5] int64

	// var newArray [3] int32 = [3]int32{1, 2, 3} OR

	// newArray := [3]int32{1, 2, 3}

	// fmt.Println(newArray[0])

	// for i:= range(newArray) {
	// 	fmt.Println(newArray[i])
	// }

	// 2D arrays


	nums := [2][3]int64{{1, 2}, {3, 4}}

	for i:= range(nums) {
		for j:= range(nums){
			fmt.Println(nums[i][j])
		}
	}
	// for i := 0; i < len(nums); i++ {
	// 	for j := 0; j < len(nums); j++ {
	// 		fmt.Println(nums[i][j])
	// 	}
		
	// }

	// // fmt.Printf(intArr[1])
	// fmt.Print(intArr)

	// // fmt.Println("Hello")


	// // Slice

	// var intSlice []int32 = []int32{1, 2, 3}

	// fmt.Println(intSlice)

	// intSlice = append(intSlice, 5)
	// fmt.Println(intSlice)


}