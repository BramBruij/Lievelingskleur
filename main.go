package main

import (
	"fmt"
	"os"
)

func main() {
	// Vraag de gebruiker voor een kleur
	fmt.Print("Wat is jouw lievelingskleur? ")
	var favoriteColor string
	fmt.Scanln(&favoriteColor)

	// Geef kleuren aan voor verschillende kleuren
	// map[string]string{ dit is de aanduiding van een map. Het geeft aan dat gedichten een map is waarvan de sleutels en waarden beide sting zijn.
	gedichten := map[string]string{
		"rood":   "Rood met passie.",
		"blauw":  "Blauw zoals de lucht.",
		"groen":  "Groen van de natuur.",
		"geel":   "Geel als de stralen van de zon.",
		"oranje": "Oranje de kleur van de zon",
		"roze":   "Roze is de kleur van een verse vrucht",
	}

	// Controlleer of de kleur in de lijst voorkomt
	gedicht, found := gedichten[favoriteColor]
	if !found {
		fmt.Println("Sorry, we hebben geen gedicht voor die kleur.")
		return
	}

	// Open een text bestand om het gedicht naartoe te zenden
	file, err := os.Create("gedicht.txt")
	if err != nil {
		fmt.Println("Fout bij het maken van het tekstbestand:", err)
		return
	}

	// Schrijf het gedicht naar het text file
	_, err = file.WriteString(gedicht)
	if err != nil {
		fmt.Println("Fout bij het schrijven naar het tekstbestand:", err)
		return
	}

	fmt.Println("Gedicht is geschreven naar gedicht.txt!")
}
