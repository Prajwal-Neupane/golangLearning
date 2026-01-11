package loops

import "fmt"



func Loops() {

	sum:= 0;

	for i := range 10 {
		sum = sum + i;
		
	}

	fmt.Printf("%v", sum);


	// for i := 0; i < count; i++ {  this can be replaced with range
		
	// }
}