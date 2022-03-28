package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 将接口编码为 JSON

type Customer map[string]interface{}
type Name map[string]interface{}
type Address map[string]interface{}

func main() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")

	john := Customer{
		"Name": Name{
			"FirstName": "John",
			"LastName":  "Smith",
		},
		"Email": "johnsmith@example.com",
		"Address": Address{
			"Line1": "The White House",
			"Line2": "1600 Pennsylvania Avenue NW",
			"Line3": "Washington, DC 20500",
		},
		"DOB": dob,
	}

	johnJSON, err := json.MarshalIndent(john, "", "   ")
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}
