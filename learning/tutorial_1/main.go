package main

import "fmt"


func main() {
	fmt.Println("Hello")

	var printValue  string = "Prajwal"
	printName(printValue)

	var numerator int = 101
	var denominator int = 20

	var result, remainder = intDivision(numerator, denominator)
	fmt.Println(result,remainder)

}

func printName(printValue string) {
	fmt.Println(printValue)

}

func intDivision(numerator int, denominator int) (int , int) {
	result := numerator/denominator
	remainder := numerator % denominator

	// Or we can also do this
	// var remainder int = numerator % denominator
	// var result int = numerator / denominator

	return result, remainder
}
