package main

import (
	"encoding/json"
	"fmt"
)

type world struct {
	Name string `json:"name"`
	Land string `json:"land"`
	Water string `json:"water,omitempty"`
	Comment string `json:"comment"`
}

func main() {
	a := world{
		"Planet X",
		"18 Continents",
		"0 litres",
		"Dead world",
	}

	b := &world{
		Name: "Planet Y",
		Land: "10 Continents",
		Comment: "Alive world",
	}

	aJson, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	bJson, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(aJson), string(bJson))
}