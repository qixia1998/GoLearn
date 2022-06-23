package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/* create a map */
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["Pakistan"] = "Islamabad"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not */
	capital, ok := countryCapitalMap["India"]

	/* if ok is true, entry is present othervise entry is absent */
	if(ok) {
		fmt.Println("Capital of India is", capital)
	} else {
		fmt.Println("Capital of India is not present")
	}
}