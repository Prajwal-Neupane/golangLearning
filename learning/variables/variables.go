package variables

import "fmt"


func Variables() {

	// Method 1

	var age int = 17
	var name string = "Prajwal"
	var isStudent bool = false

	// Method 2

	age1 := 18
	name1 := "Messi"
	isStudent1 := true


	// Multiple Variables

	var a,b,c = 1, 2, 3
	x, y, z := "Prajwal", "Ram", "god"

	fmt.Println(age, age1, name, name1, isStudent, isStudent1, a, b, c, x, y, z)

	const pi = 3.14

	fmt.Println(pi)



}