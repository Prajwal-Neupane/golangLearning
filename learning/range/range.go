package main

import "fmt"

func main() {

newSlice := make([]int, 3)

newSlice = append(newSlice, 1, 2, 3)


sum:= 0

for _, newSlices:= range newSlice {
	sum = sum + newSlices
} 

for i, num:= range newSlice {
	sum += num

	fmt.Println(i, sum)
}

fmt.Println(sum)


newMap := map[string]int{"roll1":1, "roll2": 2}

for i, mapValues:= range newMap {
	fmt.Println(i, mapValues)
}

for i, c:= range "golang" { 
	fmt.Println(i,c) //just printing "c" will print the uni-code value of the words "golang"
}

for i,c:= range "golang" {
	fmt.Println(i, string(c))  //using string(c) will print the actual value "golang"
}


}
