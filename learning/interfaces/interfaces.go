package main

import "fmt"

type payment struct {

}

func (p payment) makePayment(amount float32) {
	esewaPaymentGateway := esewa{}
	esewaPaymentGateway.pay(amount)

}


type esewa struct {

}

func (e esewa) pay(amount float32) {
	// Logic to make payment

	fmt.Println("Making payment through esewa", amount)
}


func main() {

	newPayment := payment{}
	newPayment.makePayment(100)

}