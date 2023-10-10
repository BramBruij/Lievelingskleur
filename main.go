package main

import (
	"fmt"
)

func main() {
	// Vraag de gebruiker voor een kleur
	fmt.Print("Wat is jouw lievelingskleur? ")
	var favoriteColor string
	fmt.Scanln(&favoriteColor)

	// Geef kleuren aan voor verschillende kleuren
	gedicht := map[string]string{
		"rood":  "Rood met passie.",
		"blauw": "Blauw zoals de lucht.",
		"groen": "Groen van de natuur.",
		"geel":  "Geel als de stralen van de zon.",
	}

	// Controlleer of de kleur in de lijst voorkomt
	gedicht, found := gedicht[favoriteColor]
	if !found {
		fmt.Println("Sorry, we hebben geen gedicht voor die kleur.")
		return
	}
}
