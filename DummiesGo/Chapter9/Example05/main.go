package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString :=
		`{
            "success": true,
            "timestamp": 1588779306,
            "base": "EUR",
            "date": "2020-05-06",
            "rates": {
                "AUD": 1.683349,
                "CAD": 1.528643,
                "GBP": 0.874757,
                "SGD": 1.534513,
                "USD": 1.080054
            }
        }`

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString), &result)
	fmt.Println(result["success"])

	rates := result["rates"]
	fmt.Println(rates)

	currencies := rates.(map[string]interface{})
	SGD := currencies["SGD"]
	fmt.Println(SGD)
}
