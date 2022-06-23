package main

import "fmt"

func main() {
	/* create a map */
	countryCapitaMap := map[string]string {"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "Pakistan": "Islamabad"}

	fmt.Println("Original map")

	/* print map */
	for country := range countryCapitaMap {
		fmt.Println("Capital of", country, "is", countryCapitaMap[country])
	}

	/* delete an entry */
	delete(countryCapitaMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	/* print map */
	for country := range countryCapitaMap {
		fmt.Println("Capital of", country, "is", countryCapitaMap[country])
	}
}