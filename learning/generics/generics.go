package main

import "fmt"

func printSlice [T any](items [] T) {
	for _, item :=range items {
		fmt.Println(item)
	}
}

func printNames (items [] string) {
	for _, item:= range items {
		fmt.Println(item)
	}
}

func main () {
	// nums:= []int {1, 2, 3}
	names:= []string {"prajwal", "rohan", "rupesh"}
	printSlice(names)
	// printNames(names)


}