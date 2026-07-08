package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Übung 1: Personen als JSON speichern und lesen
	fmt.Println("=== Übung 1: JSON ===")
	personen := []map[string]interface{}{
		{"name": "Anna", "alter": 28},
		{"name": "Ben", "alter": 34},
	}
	daten, _ := json.MarshalIndent(personen, "", "  ")
	os.WriteFile("personen.json", daten, 0644)
	fmt.Println("personen.json geschrieben")

	gelesen, _ := os.ReadFile("personen.json")
	var geladen []map[string]interface{}
	json.Unmarshal(gelesen, &geladen)
	for _, p := range geladen {
		fmt.Printf("  %v (%v)\n", p["name"], p["alter"])
	}
	os.Remove("personen.json")

	// Übung 2: CSV mit Noten, Durchschnitt berechnen
	fmt.Println("\n=== Übung 2: CSV Noten ===")
	csvText := `Name,Fach,Note
Anna,Mathe,1.5
Anna,Deutsch,2.0
Ben,Mathe,2.5
Ben,Deutsch,1.0`
	reader := csv.NewReader(strings.NewReader(csvText))
	records, _ := reader.ReadAll()
	notenProSchueler := make(map[string][]float64)
	for _, row := range records[1:] {
		note := 0.0
		fmt.Sscanf(row[2], "%f", &note)
		notenProSchueler[row[0]] = append(notenProSchueler[row[0]], note)
	}
	for name, noten := range notenProSchueler {
		sum := 0.0
		for _, n := range noten {
			sum += n
		}
		fmt.Printf("  %s: Durchschnitt %.1f\n", name, sum/float64(len(noten)))
	}

	// Übung 3: Datei Zeile für Zeile mit Zeilennummer
	fmt.Println("\n=== Übung 3: Zeile für Zeile ===")
	os.WriteFile("testdatei.txt", []byte("Erste\nZweite\nDritte\n"), 0644)
	datei, _ := os.Open("testdatei.txt")
	defer datei.Close()
	scanner := bufio.NewScanner(datei)
	nummer := 1
	for scanner.Scan() {
		fmt.Printf("  %d: %s\n", nummer, scanner.Text())
		nummer++
	}
	os.Remove("testdatei.txt")

	// Übung 4: Notiz-Speicher
	fmt.Println("\n=== Übung 4: Notiz-Speicher ===")
	notizen := []string{"Go lernen", "Übungen machen", "Kaffee holen"}
	notizDaten, _ := json.MarshalIndent(notizen, "", "  ")
	os.WriteFile("notizen.json", notizDaten, 0644)
	fmt.Println("notizen.json geschrieben")
	geladenNotiz, _ := os.ReadFile("notizen.json")
	var ausgabe []string
	json.Unmarshal(geladenNotiz, &ausgabe)
	for i, n := range ausgabe {
		fmt.Printf("  %d: %s\n", i+1, n)
	}
	os.Remove("notizen.json")
}
