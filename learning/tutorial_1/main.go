package main

import (
	"errors"
	"fmt"
)


func main() {
	fmt.Println("Hello")

	var printValue  string = "Prajwal"
	printName(printValue)

	var numerator int = 101
	var denominator int = 0

	var result, remainder, error = intDivision(numerator, denominator)

	if error != nil {
		fmt.Printf(error.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {

		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

}

func printName(printValue string) {
	fmt.Println(printValue)

}

func intDivision(numerator int, denominator int) (int , int, error) {
	var err error

	if denominator == 0{
		err = errors.New("Cannot divide by 0")
		return 0 , 0 , err
	}
	result := numerator/denominator
	remainder := numerator % denominator

	// Or we can also do this
	// var remainder int = numerator % denominator
	// var result int = numerator / denominator

	return result, remainder, err
}
