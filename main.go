package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func main() {
	fmt.Println("Lade aktuelle Wechselkurse von der API...")
	ratesData, err := FetchRates() // Nutzt die Logik aus api.go
	if err != nil {
		log.Fatalf("Fehler beim Laden der Kurse: %v", err)
	}
	fmt.Println("Kurse erfolgreich geladen!\n")

	var amountStr string
	var from string
	var to string

	// Interaktives Formular mit github.com/charmbracelet/huh erstellen
	form := huh.NewForm(
		huh.NewGroup(
			// 1. Betrag eingeben
			huh.NewInput().
				Title("Betrag").
				Placeholder("100").
				Value(&amountStr),

			// 2. Ausgangswährung auswählen
			huh.NewSelect[string]().
				Title("Ausgangswährung").
				Options(
					huh.NewOption("US Dollar (USD)", "USD"),
					huh.NewOption("Euro (EUR)", "EUR"),
					huh.NewOption("Britisches Pfund (GBP)", "GBP"),
					huh.NewOption("Japanischer Yen (JPY)", "JPY"),
				).
				Value(&from),

			// 3. Zielwährung auswählen
			huh.NewSelect[string]().
				Title("Zielwährung").
				Options(
					huh.NewOption("US Dollar (USD)", "USD"),
					huh.NewOption("Euro (EUR)", "EUR"),
					huh.NewOption("Britisches Pfund (GBP)", "GBP"),
					huh.NewOption("Japanischer Yen (JPY)", "JPY"),
				).
				Value(&to),
		),
	)

	// Formular ausführen
	err = form.Run()
	if err != nil {
		log.Fatalf("Fehler im Formular: %v", err)
	}

	// String-Eingabe in Float konvertieren
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Fatalf("Ungültiger Betrag: %v", err)
	}

	// Umrechnung über converter.go durchführen[cite: 2]
	result, err := ConvertCurrency(amount, from, to, ratesData.Rates)
	if err != nil {
		log.Fatalf("Umrechnungsfehler: %v", err)
	}

	// Schöne Ausgabe des Ergebnisses
	fmt.Printf("\n✨ Ergebnis: %.2f %s = %.2f %s\n", amount, from, result, to)
}
