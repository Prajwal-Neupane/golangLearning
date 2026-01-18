package main

import "fmt"


type OrderStatus int

const(
	Received OrderStatus = iota
	Delivered
	Processed
	Prepared 
	Confirmed
)


func changeOrderStatus (staus OrderStatus) {
	fmt.Println("Order status changed to", staus)
}
func main() {

	changeOrderStatus(Received)

}