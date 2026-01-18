package main

import "fmt"


type paymenter interface {
	pay(amount float32)   //interface of paymenter which has pay method which will receive amount variable of float32 type
}

type payment struct {

	gateway paymenter
}





func (p payment) makePayment(amount float32) {
	// esewaPaymentGateway := esewa{}
	// esewaPaymentGateway.pay(amount)
	// khaltiPaymentGateway := khalti{}

	p.gateway.pay(amount)
}


type esewa struct {

}



func (e esewa) pay(amount float32) {
	// Logic to make payment

	fmt.Println("Making payment through esewa", amount)
}

type khalti struct {

}

func (k khalti) pay (amount float32) {

	fmt.Println("Making payment through khalti", amount)

}

type paypal struct {

}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment through paypal")
}


func main() {

	// newPayment := payment{}
	// newPayment.makePayment(100)

	// khaltiPaymentGateway := khalti{}
	paypalPaymentGateway := paypal{}

	newPayment := payment{
		// gateway: khaltiPaymentGateway,
		gateway: paypalPaymentGateway,
	}	// newPayment.makePayment(100)

	newPayment.makePayment(100)
}