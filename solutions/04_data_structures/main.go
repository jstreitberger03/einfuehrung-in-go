package main

import "fmt"

func main() {
	// Übung 1: Array mit 3 Lieblingszahlen
	lieblingszahlen := [3]int{7, 13, 42}
	fmt.Println("Lieblingszahlen:", lieblingszahlen)

	// Übung 2: Leeres Slice, 5 Elemente hinzufügen
	var zahlen []int
	for i := 1; i <= 5; i++ {
		zahlen = append(zahlen, i*10)
	}
	fmt.Println("Slice:", zahlen)

	// Übung 3: Summe aller Zahlen in einem Slice
	summe := 0
	for _, z := range zahlen {
		summe += z
	}
	fmt.Println("Summe:", summe)

	// Übung 4: Map mit 5 Städten und Einwohnerzahlen
	stadtEinwohner := map[string]int{
		"Wien":    1_950_000,
		"Berlin":  3_700_000,
		"München": 1_500_000,
		"Hamburg": 1_900_000,
		"Zürich":  420_000,
	}
	for stadt, einwohner := range stadtEinwohner {
		fmt.Printf("  %s: %d Einwohner\n", stadt, einwohner)
	}

	// Übung 5: Prüfen, ob ein Wert im Slice existiert
	suche := 30
	gefunden := contains(zahlen, suche)
	fmt.Printf("\n%d im Slice? %v\n", suche, gefunden)

	// Übung 6: Buchstaben-Zähler
	text := "Hello World"
	zaehler := make(map[rune]int)
	for _, ch := range text {
		zaehler[ch]++
	}
	fmt.Println("\nBuchstaben-Zähler in \"Hello World\":")
	for ch, anzahl := range zaehler {
		fmt.Printf("  '%c': %d\n", ch, anzahl)
	}
}

// contains prüft, ob ein Wert im Slice existiert
func contains(slice []int, wert int) bool {
	for _, v := range slice {
		if v == wert {
			return true
		}
	}
	return false
}
