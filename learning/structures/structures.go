package structures

import "fmt"

type gasEngine struct {
	liters uint8
	kmpl uint16
}


func main() {

	var myEngine gasEngine = gasEngine{liters: 100, kmpl: 20}
	// myEngine.kmpl = 100

	fmt.Println(myEngine.kmpl, myEngine.liters)


}