package main

import "fmt"

func main() {
	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	/* create a map */
	countryCapitaMap := map[string]string {"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys */
	for country := range countryCapitaMap {
		fmt.Println("Capital of", country, "is", countryCapitaMap[country])
	}

	/* print map using key-value */
	for country, capital := range countryCapitaMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
