package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {
	stripe.Key = "pk_test_51NvCfECu6GjgOc9s24qvz2y1zR0QLy2ob1KYJRZOz74KnjmK2SCsxtREzJQRfAsL5x1tqX45LSL345HWBujbinMT00BfLXu2ah"
	params := &stripe.CustomerParams{}
	params.SetStripeAccount("acct_1032D82eZvKYlo2C")
	_, err := customer.Get("cu_19YMK02eZvKYlo2CYWjsbgL3", params)

	if err != nil {
		// Try to safely cast a generic error to a stripe.Error so that we can get at
		// some additional Stripe-specific information about what went wrong.
		if stripeErr, ok := err.(*stripe.Error); ok {
			// The Code field will contain a basic identifier for the failure.
			switch stripeErr.Code {
			case stripe.ErrorCodeCardDeclined:
			case stripe.ErrorCodeExpiredCard:
			case stripe.ErrorCodeIncorrectCVC:
			case stripe.ErrorCodeIncorrectZip:
				// etc.
			}

			// The Err field can be coerced to a more specific error type with a type
			// assertion. This technique can be used to get more specialized
			// information for certain errors.
			if cardErr, ok := stripeErr.Err.(*stripe.CardError); ok {
				fmt.Printf("Card was declined with code: %v\n", cardErr.DeclineCode)
			} else {
				fmt.Printf("Other Stripe error occurred: %v\n", stripeErr.Error())
			}
		} else {
			fmt.Printf("Other error occurred: %v\n", err.Error())
		}
	}

}
