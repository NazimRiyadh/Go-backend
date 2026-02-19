package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (r *payment) makepayment(amount float32) {
	r.gateway.pay(amount)
}

type razorpay struct {
}

func (r *razorpay) pay(amount float32) {
	fmt.Println("Payment done by Razorpay ", amount)
}

type stripe struct {
}

func (r *stripe) pay(amount float32) {
	fmt.Println("Payment done by STRIPE ", amount)
}

func main() {
	//stripe := stripe{}

	razorpay := razorpay{}
	payment := payment{
		gateway: &razorpay,
	}
	payment.makepayment(100)
}
