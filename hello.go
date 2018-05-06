package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Japan": "Tokyo",
		"China": "Perking"}

	for country := range countryCapitalMap {
		fmt.Println(country, "Cap", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")

	fmt.Println("F Gone")

	for country := range countryCapitalMap {
		fmt.Println(country, "Cap", countryCapitalMap[country])
	}
}
