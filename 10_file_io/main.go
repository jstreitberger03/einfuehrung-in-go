// Kapitel 10: Dateien & Daten — Lesen und Schreiben
//
// 🎯 Lernziele:
// - Dateien lesen und schreiben
// - JSON-Daten verarbeiten (encoding/json)
// - CSV-Daten lesen
// - io.Reader und io.Writer verstehen
//
// 📖 Los geht's!

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ==========================================
	// 1. Dateien schreiben
	// ==========================================
	fmt.Println("=== Dateien schreiben ===")

	// os.WriteFile schreibt einen kompletten String in eine Datei
	text := []byte("Hallo Datei!\nDas ist eine Zeile.\nUnd noch eine.\n")
	err := os.WriteFile("beispiel.txt", text, 0644)
	if err != nil {
		fmt.Println("Fehler beim Schreiben:", err)
		return
	}
	fmt.Println("✅ beispiel.txt wurde geschrieben")

	// ==========================================
	// 2. Dateien lesen
	// ==========================================
	fmt.Println("\n=== Dateien lesen ===")

	daten, err := os.ReadFile("beispiel.txt")
	if err != nil {
		fmt.Println("Fehler beim Lesen:", err)
		return
	}
	fmt.Println("Inhalt von beispiel.txt:")
	fmt.Println(string(daten))

	// ==========================================
	// 3. JSON — Daten strukturiert speichern
	// ==========================================
	fmt.Println("\n=== JSON ===")

	// Strukturierte Daten in JSON speichern
	person := PersonJSON{
		Name:   "Anna",
		Alter:  28,
		Stadt:  "Wien",
		Hobbys: []string{"Lesen", "Wandern", "Kochen"},
		Aktiv:  true,
	}

	// Go → JSON (Marshal)
	jsonDaten, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Fehler beim JSON-Erzeugen:", err)
		return
	}
	fmt.Println("JSON-Ausgabe:")
	fmt.Println(string(jsonDaten))

	// JSON in Datei schreiben
	os.WriteFile("person.json", jsonDaten, 0644)
	fmt.Println("✅ person.json wurde geschrieben")

	// JSON aus Datei lesen und parsen
	jsonGelesen, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Println("Fehler beim Lesen:", err)
		return
	}

	var geladenePerson PersonJSON
	err = json.Unmarshal(jsonGelesen, &geladenePerson)
	if err != nil {
		fmt.Println("Fehler beim JSON-Parsen:", err)
		return
	}
	fmt.Printf("\nGeladene Person: %s (%d) aus %s\n",
		geladenePerson.Name, geladenePerson.Alter, geladenePerson.Stadt)

	// ==========================================
	// 4. CSV lesen
	// ==========================================
	fmt.Println("\n=== CSV ===")

	csvText := `Name,Alter,Stadt
Anna,28,Wien
Ben,34,München
Clara,22,Hamburg`

	reader := csv.NewReader(strings.NewReader(csvText))
	datensaetze, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Fehler beim CSV-Lesen:", err)
		return
	}

	fmt.Println("CSV-Tabelle:")
	for _, zeile := range datensaetze {
		fmt.Printf("%-15s %-5s %-15s\n", zeile[0], zeile[1], zeile[2])
	}

	// ==========================================
	// 5. Weitere nützliche Datei-Operationen
	// ==========================================
	fmt.Println("\n=== Weitere Operationen ===")

	// Prüfen, ob eine Datei existiert
	info, err := os.Stat("go.mod")
	if err == nil {
		fmt.Printf("✅ go.mod existiert (%d Bytes)\n", info.Size())
	} else if os.IsNotExist(err) {
		fmt.Println("❌ go.mod existiert nicht")
	}

	// Datei löschen
	os.Remove("beispiel.txt")
	os.Remove("person.json")
	fmt.Println("✅ Temporäre Dateien gelöscht")

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Struktur für JSON ---

type PersonJSON struct {
	Name   string   `json:"name"`
	Alter  int      `json:"alter"`
	Stadt  string   `json:"stadt"`
	Hobbys []string `json:"hobbys"`
	Aktiv  bool     `json:"aktiv"`
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Schreibe eine Liste von Personen als JSON in eine Datei und lies sie zurück
// 2. Lese eine CSV-Datei mit Noten (Name, Fach, Note) und berechne den Durchschnitt
// 3. Schreibe eine Funktion, die eine Datei Zeile für Zeile einliest und
//    jede Zeile mit Zeilennummer ausgibt
// 4. Erstelle einen einfachen Notiz-Speicher: Notizen in JSON-Datei speichern
// ---------------------------------------------------------------------------
