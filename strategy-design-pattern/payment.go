package main

import "fmt"

type PaymentStrategy interface {
	ProcessPayment()
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment() {
	fmt.Println("Process credit card payment...")
}

type PayPalPayment struct{}

func (p *PayPalPayment) ProcessPayment() {
	fmt.Println("Process PayPal payment...")
}

type CryptoPayment struct{}

func (c *CryptoPayment) ProcessPayment() {
	fmt.Println("Process crypto payment...")
}

type StripePayment struct{}

func (s *StripePayment) ProcessPayment() {
	fmt.Println("Process Stripe payment...")
}

type PaymentProcessor struct {
	paymentStrategy PaymentStrategy
}

func NewPaymentProcessor(paymentStrategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{paymentStrategy}
}

func (pp *PaymentProcessor) ProcessPayment() {
	pp.paymentStrategy.ProcessPayment()
}

func (pp *PaymentProcessor) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	pp.paymentStrategy = paymentStrategy
}
