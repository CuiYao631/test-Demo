package main

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
)

func main() {
	stripe.Key = "sk_test_51NvCfECu6GjgOc9scJiWFP0YHrum0pVcCyJdyA02BoDrC8O7xN3D1YBLRQ7K4l2mWp3Jq2qjhzkKB4yAK6U0HNlQ002qfgUILm"
	//设置产品参数
	product_params := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	//创建一个产品
	starter_product, _ := product.New(product_params)

	//设置价格参数
	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	//创建一个价格
	starter_price, _ := price.New(price_params)

	fmt.Println("Success! Here is your starter subscription product id: " + starter_product.ID)
	fmt.Println("Success! Here is your starter subscription price id: " + starter_price.ID)
}
