package main

import "fmt"

type PaymentChargeRequest interface {
	helloFromCharge()
}
type PaymentCancelRequest interface {
	helloFromCancel()
}

type PaymentFactory interface {
	createPaymentChargeRequest() PaymentChargeRequest
	createPaymentCancelRequest() PaymentCancelRequest
}

type CreditCard struct {
	ID string
}

func (c *CreditCard) helloFromCharge() {
	fmt.Println("hello from charge")
}

func (c *CreditCard) helloFromCancel() {
	fmt.Println("hello from cancel")
}

func (c *CreditCard) createPaymentChargeRequest() PaymentChargeRequest {
	return c
}

func (c *CreditCard) createPaymentCancelRequest() PaymentCancelRequest {
	return c
}

func main() {
	card := CreditCard{}
	charge := card.createPaymentChargeRequest()
	cancel := card.createPaymentCancelRequest()

	charge.helloFromCharge()
	cancel.helloFromCancel()
}
