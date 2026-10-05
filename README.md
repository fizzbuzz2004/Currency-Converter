# Go Currency Converter

Ein moderner Kommandozeilen-Währungsrechner (TUI) geschrieben in Go, der Live-Wechselkurse über die Open Exchange Rates API abruft und ein interaktives Formular über das Charm Bracelet `huh`-Paket bereitstellt.

## Projektstruktur

* **`api.go`**: Zuständig für die Kommunikation mit der Open Exchange Rates API und das Laden der JSON-Daten.
* **`converter.go`**: Enthält die mathematische Logik zur Umrechnung der Währungen auf Basis von USD.
* **`main.go`**: Die Hauptdatei, die ein interaktives TUI-Formular (`github.com/charmbracelet/huh`) für die Benutzereingabe bereitstellt.
* **`go.mod`**: Definiert das Go-Modul (`currency-converter`).

## Verwendete Pakete

* `net/http`: Für HTTP-Anfragen an die Währungs-API.
* `encoding/json`: Zum Parsen und Verarbeiten der API-Daten.
* `github.com/charmbracelet/huh`: Für die interaktive TUI-Benutzeroberfläche im Terminal.

## Unterstützte Währungen

Die Anwendung unterstützt standardmäßig folgende Hauptwährungen über das Auswahlmenü:
* **USD** (US Dollar)
* **EUR** (Euro)
* **GBP** (Britisches Pfund)
* **JPY** (Japanischer Yen)

## Voraussetzungen

* Go (installierte Version 1.22 oder neuer)
* Ein kostenloser Account bei [Open Exchange Rates](https://openexchangerates.org/) für eine App ID.

## Installation & Ausrichtung

1. Klone oder öffne das Projekt in deinem Terminal im Projektverzeichnis.

2. Installiere das TUI-Paket von Charm Bracelet:
   ```bash
   go get [github.com/charmbracelet/huh](https://github.com/charmbracelet/huh)
