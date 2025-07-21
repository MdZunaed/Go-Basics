package main

import "fmt"

func main() {
	newPayment := payment{
		gateway: paypal{},
		// gateway: bkash{},
		// gateway: testPayment{},
	}
	newPayment.makePayment(100)
}

type paymenter interface {
	pay(ammount float32) // Must be implement this method to work
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(ammount float32) {
	//bkashGateway:= bkash{}
	p.gateway.pay(ammount)
}

type bkash struct{}

func (p bkash) pay(ammount float32) {
	fmt.Println("Making payment using bkash", ammount)
}

type paypal struct{}

func (p paypal) pay(ammount float32) {
	fmt.Println("Making payment using paypal", ammount)
}

type testPayment struct{}

func (p testPayment) pay(ammount float32) {
	fmt.Println("Making payment using TestPay", ammount)
}
