package datatypes

import "fmt"

func DataTypes() {

	// Constants
	const name = "Prajwal"
	const age = 13

	// Multiple constants

	const (
		A=1 
	 	B=2
	    C=3
		
		)

		// fmt.Println(A, B, C)


		// Type Conversion

		a := 5
		var b float32 = 2.5

		//var c = a + b  //It will cause error because of type mismatched
		var c = float32(a) + b

		fmt.Println(c)


}