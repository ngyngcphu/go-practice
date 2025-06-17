package main

func main() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}
	crypto := &CryptoPayment{}
	stripe := &StripePayment{}

	processor := NewPaymentProcessor(creditCard)
	processor.ProcessPayment()

	processor.SetPaymentStrategy(payPal)
	processor.ProcessPayment()

	processor.SetPaymentStrategy(crypto)
	processor.ProcessPayment()

	processor.SetPaymentStrategy(stripe)
	processor.ProcessPayment()
}
