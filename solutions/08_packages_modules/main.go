package main

import "fmt"

// Übung 1: Paket "wetter" mit TemperaturBeschreibung
// (In einem echten Projekt würde man einen Unterordner "wetter/" erstellen.
//
//	Hier zeigen wir die Struktur inline.)
type wetterService struct{}

func (w wetterService) TemperaturBeschreibung(temp float64) string {
	switch {
	case temp < 0:
		return "Es ist unterkühlt"
	case temp < 10:
		return "Es ist kalt"
	case temp < 20:
		return "Es ist kühl"
	case temp < 30:
		return "Es ist angenehm"
	default:
		return "Es ist heiß"
	}
}

// Übung 2: Nicht-exportierte Hilfsfunktion
// (In taschenrechner.go hinzufügen:)
// func quadrat(x int) int { return x * x } // klein = nicht exportiert

// Übung 3: Externes Paket (hier nur demonstriert, go get nötig)
// go get github.com/fatih/color
// import "github.com/fatih/color"
// color.Red("Fehler!")
// color.Green("Erfolg!")

// Übung 4: init()-Funktion
// func init() { fmt.Println("Paket wurde geladen!") }
// Wird automatisch aufgerufen, bevor main() startet.

func main() {
	// Übung 1
	w := wetterService{}
	for _, temp := range []float64{-5, 5, 15, 25, 35} {
		fmt.Printf("%.0f°C: %s\n", temp, w.TemperaturBeschreibung(temp))
	}

	// Übung 2-4: Siehe Kommentare oben.
	// Erstelle die Dateien in einem echten Unterordner, um sie zu testen.
	fmt.Println("\nÜbungen 2-4 erfordern echte Pakete. Siehe Kommentare im Code.")
}
