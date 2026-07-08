package main

import (
	"fmt"
	"strings"
)

func main() {
	// Übung 1: swap mit Pointern
	a, b := 5, 10
	fmt.Printf("Vor swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("Nach swap: a=%d, b=%d\n", a, b)

	// Übung 2: String in Großbuchstaben umwandeln
	text := "hallo go welt"
	toUpper(&text)
	fmt.Println("\nGroßbuchstaben:", text)

	// Übung 3: Bankkonto mit Pointer-Empfänger
	konto := Bankkonto{Inhaber: "Anna", Kontostand: 1000.0}
	fmt.Printf("\n%s: %.2f €\n", konto.Inhaber, konto.Kontostand)
	konto.einzahlen(500.0)
	fmt.Printf("Nach Einzahlung: %.2f €\n", konto.Kontostand)
	konto.abheben(200.0)
	fmt.Printf("Nach Abhebung: %.2f €\n", konto.Kontostand)
}

// Übung 1: swap vertauscht zwei Werte über Pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Übung 2: toUpper wandelt einen String in Großbuchstaben um
func toUpper(s *string) {
	*s = strings.ToUpper(*s)
}

// Übung 3: Bankkonto mit Pointer-Empfängern
type Bankkonto struct {
	Inhaber    string
	Kontostand float64
}

func (k *Bankkonto) einzahlen(betrag float64) {
	k.Kontostand += betrag
}

func (k *Bankkonto) abheben(betrag float64) {
	if betrag > k.Kontostand {
		fmt.Println("Nicht genug Guthaben!")
		return
	}
	k.Kontostand -= betrag
}
