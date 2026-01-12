package main

import "fmt"

func addNumbers (a ,b int ) int{
	return a + b

}

func returnString () (string, string , string) {
	return "string1", "string2", "string3"
}
func main() {


	result := addNumbers(5, 2)
	fmt.Println(result)

	fmt.Println(returnString())

	string1, string2, string3 := returnString()
	fmt.Println(string1, string2, string3)


}