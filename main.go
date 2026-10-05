package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Kurse einmalig beim Start laden
	fmt.Println("Lade aktuelle Wechselkurse von der API...")
	ratesData, err := FetchRates()
	if err != nil {
		log.Fatalf("Fehler beim Laden der Kurse: %v", err)
	}
	fmt.Println("Kurse erfolgreich geladen!\n-----------------------------------")

	reader := bufio.NewReader(os.Stdin)

	for {
		// 1. Betrag abfragen
		fmt.Print("Gib den Betrag ein (oder 'q' zum Beenden): ")
		amountInput, _ := reader.ReadString('\n')
		amountInput = strings.TrimSpace(amountInput)

		if strings.ToLower(amountInput) == "q" {
			fmt.Println("Programm beendet. Tschüss!")
			break
		}

		amount, err := strconv.ParseFloat(amountInput, 64)
		if err != nil {
			fmt.Println("❌ Ungültiger Betrag. Bitte gib eine Zahl ein.\n")
			continue
		}

		// 2. Ausgangswährung abfragen
		fmt.Print("Ausgangswährung (z.B. USD, EUR, GBP): ")
		fromInput, _ := reader.ReadString('\n')
		from := strings.ToUpper(strings.TrimSpace(fromInput))

		// 3. Zielwährung abfragen
		fmt.Print("Zielwährung (z.B. USD, EUR, GBP): ")
		toInput, _ := reader.ReadString('\n')
		to := strings.ToUpper(strings.TrimSpace(toInput))

		// 4. Umrechnung berechnen
		result, err := ConvertCurrency(amount, from, to, ratesData.Rates)
		if err != nil {
			fmt.Printf("❌ Fehler: %v\n\n", err)
			continue
		}

		// 5. Ergebnis ausgeben
		fmt.Printf("✅ Ergebnis: %.2f %s = %.2f %s\n", amount, from, result, to)
		fmt.Println("-----------------------------------")
	}
}
