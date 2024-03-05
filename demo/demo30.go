package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	//商户代码
	MerchantCode string `json:"merchant_code"`
	//账单代码
	BillCode string `json:"biller_code"`
	//商家订单号
	MerchantOrderNo string `json:"merchant_order_no"`
	//买家id
	BuyerId string `json:"buyer_id"`
	//用户ID
	UserId string `json:"user_id"`
	//货币
	Currency string `json:"currency"`
	//交易类型
	TradeType string `json:"trade_type"`
	//订单金额
	OrderAmount int `json:"order_amount"`
	//订单描述
	OrderDescription string `json:"order_description"`
	//商品信息
	GoodsInfo string `json:"goods_info"`
	//返回网址
	ReturnUrl string `json:"return_url"`
	//随机数
	NonceStr string `json:"nonce_str"`
	//通知网址
	NotifyUrl string `json:"notify_url"`
	//订单超时
	OrderTimeout string `json:"order_timeout"`
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj).Elem() // 获取指向结构体值的引用
	t := v.Type()                    // 获取类型信息

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)             // 获取每个字段的类型信息
		value := v.Field(i).Interface() // 获取每个字段的值
		if field.Tag != "" && field.Tag.Get("json") != "-" {
			key := field.Tag.Get("json") // 从tag中提取JSON标记名称
			result[key] = value          // 将字段名作为键，字段值作为值存入map中
		} else if field.Anonymous { // 如果字段是匿名字段，则递归调用StructToMap函数处理该字段
			embeddedResult := StructToMap(value)
			for k, v := range embeddedResult {
				result[k] = v
			}
		}
	}

	return result
}

func main() {
	p1 := &Person{
		MerchantCode:     "UPTEST",                                                 //商户代码
		BillCode:         "2ciKVWw2xkzbyDQdUoOIQT0K6NF",                            //账单代码
		MerchantOrderNo:  "2ciKVWw2xkzbyDQdUoOIQT0K6NF",                            //商家订单号
		BuyerId:          "2Zq7KWd0IDf5o1CPWs8oILmaweI",                            //买家id
		UserId:           "2Zq7KWd0IDf5o1CPWs8oILmaweI",                            //用户ID
		Currency:         "AUD",                                                    //货币
		TradeType:        "WEB",                                                    //交易类型
		OrderAmount:      222,                                                      //订单金额
		OrderDescription: "test",                                                   //订单描述
		GoodsInfo:        "test",                                                   //商品信息
		ReturnUrl:        "https://example.com/success2ahKOuCq06F6XU2rk84cK3Y6J99", //返回网址
		NonceStr:         "cuxYcytR7O",                                             //随机数
		NotifyUrl:        "https://oeonline.com.au/webhookdev",                     //通知网址
		OrderTimeout:     "10m",                                                    //订单超时
	}
	m := StructToMap(p1)
	fmt.Println(m)
}
