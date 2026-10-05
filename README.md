# Go Currency Converter

Ein einfacher Kommandozeilen-Währungsrechner geschrieben in Go, der Live-Wechselkurse über die Open Exchange Rates API abruft.

## Projektstruktur

* **`api.go`**: Zuständig für die Kommunikation mit der Open Exchange Rates API und das Laden der JSON-Daten.
* **`converter.go`**: Enthält die mathematische Logik zur Umrechnung der Währungen auf Basis von USD.
* **`main.go`**: Die Hauptdatei mit einer interaktiven CLI-Schleife (`bufio`) für die Benutzereingabe.
* **`go.mod`**: Definiert das Go-Modul (`currency-converter`).

## Voraussetzungen

* Go (installierte Version 1.22 oder neuer)
* Ein kostenloser Account bei [Open Exchange Rates](https://openexchangerates.org/) für eine App ID.

## Installation & Ausrichtung

1. Klone oder öffne das Projekt in deinem Terminal im Projektverzeichnis.

2. Setze deine API-ID als Umgebungsvariable (`OER_APP_ID`):

   * **PowerShell (Windows):**
     ```powershell
     $env:OER_APP_ID="DEINE_API_ID_HIER"
     ```

   * **Bash / Linux / macOS:**
     ```bash
     export OER_APP_ID="DEINE_API_ID_HIER"
     ```

3. Starte das Programm:
   ```bash
   go run main.go api.go converter.go
