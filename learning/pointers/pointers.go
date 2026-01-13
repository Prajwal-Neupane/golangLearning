package main

import "fmt"



func changeNumber(a *int) { //*int we're passing reference of the number 0Xdklfds
	*a = 5;   // we are dereferencing the address with a value *a
	fmt.Println(*a)

}
func main(){

	x := 10
	// var p* x;
	// p = &x
	changeNumber(&x)
	fmt.Println(x)

}