package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
	}

	byteString, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%s\n", string(byteString))

	var film Movie
	var str string = `{"released":1942,"color,omitempty":false}`
	err = json.Unmarshal([]byte(str), &film)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("film released : %d, color,omitempty : %v\n", film.Year, film.Color)

}
