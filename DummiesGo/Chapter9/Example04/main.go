package main

// 映射自定义属性名
import (
	"encoding/json"
	"fmt"
)

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func main() {
	jsonString :=
		`{
			"base currency": "EUR",
			"destination currency": "USD"
		}`

	var rates Rates
	json.Unmarshal([]byte(jsonString), &rates)
	fmt.Println(rates.Base)
	fmt.Println(rates.Symbol)
}
