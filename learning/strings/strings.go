package strings

import "fmt"



func Strings() {

	// Strings are immutable

	// first:= "Hello"
	// second:= "World"

	// fmt.Println(first + " " + second)


	var number1 int64 = 33333
	var number2 float32 = 11

	var result = float32(number1) + number2

	fmt.Println(result)

	// Differences between printf, println and print

	// fmt.println() Prints the output in new line
	// fmt.printf() c style printf with %v(default value) %s(string) %f(float) etcs
}
