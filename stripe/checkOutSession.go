package main

import (
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"log"
)

func main() {
	stripe.Key = "sk_test_51NvCfECu6GjgOc9scJiWFP0YHrum0pVcCyJdyA02BoDrC8O7xN3D1YBLRQ7K4l2mWp3Jq2qjhzkKB4yAK6U0HNlQ002qfgUILm"

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1NyUMoCu6GjgOc9susxFqjyw"),
				Quantity: stripe.Int64(2),
			},
		},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("https://example.com/success"),
	}
	params.SetStripeAccount("acct_1NyTbMEJuAYnSEoX")
	resp, err := session.New(params)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
}
