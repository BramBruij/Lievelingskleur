package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// flag voor de kleur
	var colorPtr string
	flag.StringVar(&colorPtr, "kleur", "", "Kleur (rood, blauw, groen, geel, oranje, roze.)")
	flag.Parse()
	colorPtr = strings.ToLower(colorPtr)
	//Kleur is nu leeg dit moet nog aangepast worden. if string is empty geef error.

	// Geef kleuren aan voor verschillende kleuren
	// map[string]string{ dit is de aanduiding van een map. Het geeft aan dat gedichten een map is waarvan de sleutels en waarden beide sting zijn.
	gedichten := map[string]string{
		"rood":   "Rood met passie.",
		"blauw":  "Blauw zoals de lucht.",
		"groen":  "Groen van de natuur.",
		"geel":   "Geel als de stralen van de zon.",
		"oranje": "Oranje de kleur van de zon",
		"roze":   "Roze is de kleur van een verse vrucht",
		"paars":  "Paars is de kleur van luxe.",
	}

	// Controlleer of de kleur in de lijst voorkomt
	gekozenKleur, exists := gedichten[colorPtr]
	if !exists {
		fmt.Println("Geen geldige kleur. Kies uit rood, blauw, groen, geel, oranje, roze, paars.")
		return
	}

	// Open een text bestand om het gedicht naartoe te zenden
	fileName := "gedicht.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fout bij het maken van het tekstbestand:", err)
		return
	}
	defer file.Close()

	// Schrijf het gedicht naar het text file
	_, err = file.WriteString(gekozenKleur)
	if err != nil {
		fmt.Println("Fout bij het schrijven naar het tekstbestand:", err)
		return
	}

	fmt.Printf("Gedicht is opgeslagen in %s op basis van jouw opgegeven kleur %s. \n", fileName, colorPtr)
}
