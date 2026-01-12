package loops

import "fmt"



func Loops() {

	sum:= 0;

	for i := range 10 {
		sum = sum + i;
		
	}

	fmt.Println(sum);


	// for i := 0; i < count; i++ {  this can be replaced with range
		
	// }

	// While style loop

	i := 1;
	for i <= 10 {
		fmt.Println(i)
		i+=1;
	}
}