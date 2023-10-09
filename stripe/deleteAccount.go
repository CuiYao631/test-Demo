package main

import (
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/account"
	"log"
)

func main() {
	stripe.Key = "sk_test_51NvCfECu6GjgOc9scJiWFP0YHrum0pVcCyJdyA02BoDrC8O7xN3D1YBLRQ7K4l2mWp3Jq2qjhzkKB4yAK6U0HNlQ002qfgUILm"

	a, err := account.Del("acct_1NyTHSCWchUfK5ZH", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", a)
}
