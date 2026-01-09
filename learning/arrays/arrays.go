package arrays

import "fmt"

func Slice() {


	var intArr [5] int64

	// var newArray [3] int32 = [3]int32{1, 2, 3} OR

	newArray := [3]int32{1, 2, 3}

	fmt.Println(&newArray[0])

	// fmt.Printf(intArr[1])
	fmt.Print(intArr)

	// fmt.Println("Hello")


	// Slice

	var intSlice []int32 = []int32{1, 2, 3}

	fmt.Println(intSlice)

	intSlice = append(intSlice, 5)
	fmt.Println(intSlice)


}