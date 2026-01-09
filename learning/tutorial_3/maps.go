package main

import "fmt"


func main() {

	var myMap map[string]int64 = make(map[string]int64)
	fmt.Println(myMap)

	var myMap2 = map[string]int32{"Prajwal": 22, "Ramesh": 33}

	fmt.Println(myMap2["Prajwal"])


	var age, ok = myMap2["ramesh"]

	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}



}