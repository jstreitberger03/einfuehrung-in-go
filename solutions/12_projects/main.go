package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("=== Übung 1: Taschenrechner mit Potenz und Wurzel ===")
	fmt.Printf("2^10 = %.0f\n", potenz(2, 10))
	fmt.Printf("sqrt(144) = %.0f\n", wurzel(144))

	fmt.Println("\n=== Übung 2: /api/status ===")
	fmt.Println("(Im 12_projects/main.go einen /api/status-Endpunkt hinzufügen.)")
	fmt.Println("Beispiel-Handler:")
	fmt.Println(`  http.HandleFunc("/api/status", func(w, r) {
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
  })`)

	fmt.Println("\n=== Übung 3: Todo mit Prioritäten ===")
	liste := TodoListe{}
	liste.Hinzufuegen("Go lernen", "hoch")
	liste.Hinzufuegen("Einkaufen", "niedrig")
	liste.Hinzufuegen("Sport machen", "mittel")
	liste.Anzeigen()

	fmt.Println("\n=== Übung 4: Währungsrechner EUR→USD ===")
	kurs := 1.08
	for _, eur := range []float64{10, 50, 100} {
		fmt.Printf("  %.2f EUR = %.2f USD\n", eur, eurZuUSD(eur, kurs))
	}

	fmt.Println("\n=== Übung 5: Zahlenratespiel ===")
	zahlenratespiel()
}

// --- Übung 1: Potenz und Wurzel ---

func potenz(basis, exponent float64) float64 {
	return math.Pow(basis, exponent)
}

func wurzel(zahl float64) float64 {
	return math.Sqrt(zahl)
}

// --- Übung 3: Todo mit Prioritäten ---

type Todo struct {
	Text       string
	Prioritaet string
}

type TodoListe struct {
	Aufgaben []Todo
}

func (l *TodoListe) Hinzufuegen(text, prioritaet string) {
	l.Aufgaben = append(l.Aufgaben, Todo{Text: text, Prioritaet: prioritaet})
}

func (l *TodoListe) Anzeigen() {
	for i, t := range l.Aufgaben {
		fmt.Printf("  %d. [%s] %s\n", i+1, t.Prioritaet, t.Text)
	}
}

// --- Übung 4: Währungsrechner ---

func eurZuUSD(eur, kurs float64) float64 {
	return eur * kurs
}

// --- Übung 5: Zahlenratespiel ---

func zahlenratespiel() {
	ziel := rand.Intn(100) + 1
	versuche := []string{"50", "75", "ziel"}
	fmt.Println("Rate eine Zahl zwischen 1 und 100!")
	fmt.Println("(Simuliert mit festen Eingaben für die Demonstration)")
	for _, eingabe := range versuche {
		zahl, _ := strconv.Atoi(eingabe)
		if zahl == 0 {
			continue
		}
		switch {
		case zahl < ziel:
			fmt.Printf("  %d: Zu niedrig!\n", zahl)
		case zahl > ziel:
			fmt.Printf("  %d: Zu hoch!\n", zahl)
		default:
			fmt.Printf("  %d: Richtig! Die Zahl war %d.\n", zahl, ziel)
			return
		}
	}
	fmt.Printf("Die Zahl war %d. Versuche es nochmal!\n", ziel)
}
