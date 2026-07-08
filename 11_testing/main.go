// Kapitel 11: Testen — Qualität sichern
//
// Lernziele:
// - Unit-Tests mit Go schreiben
// - Table-Driven Tests (das Go-typische Testmuster)
// - Benchmarks messen
// - Test Coverage prüfen
//
// Los geht's!
//
// Wichtiger Hinweis: Die Tests zu diesem Kapitel befinden sich in
// main_test.go im selben Ordner. Führe sie aus mit:
// go test -v .
// go test -bench .
// go test -cover .

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Testen in Go ===")
	fmt.Println()
	fmt.Println("Tests werden in Dateien mit _test.go geschrieben.")
	fmt.Println()
	fmt.Println("Führe die Tests aus:")
	fmt.Println(" go test -v . # Alle Tests ausführlich")
	fmt.Println(" go test -bench . # Benchmarks")
	fmt.Println(" go test -cover . # Testabdeckung")
	fmt.Println(" go test -coverprofile=coverage.out .")
	fmt.Println(" go tool cover -html=coverage.out")
	fmt.Println()
	fmt.Println("Öffne main_test.go, um die Tests zu sehen!")
	fmt.Println()
	fmt.Println(" Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Funktionen, die wir testen werden ---

// Summe berechnet die Summe aller Zahlen in einem Slice
func Summe(zahlen []int) int {
	summe := 0
	for _, z := range zahlen {
		summe += z
	}
	return summe
}

// Durchschnitt berechnet den Mittelwert
func Durchschnitt(zahlen []float64) (float64, error) {
	if len(zahlen) == 0 {
		return 0, errors.New("Slice ist leer")
	}
	summe := 0.0
	for _, z := range zahlen {
		summe += z
	}
	return summe / float64(len(zahlen)), nil
}

// IstPrimzahl prüft, ob eine Zahl eine Primzahl ist
func IstPrimzahl(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci berechnet die n-te Fibonacci-Zahl
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n darf nicht negativ sein")
	}
	if n <= 1 {
		return n, nil
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// Max findet die größte Zahl in einem Slice
func Max(zahlen []int) (int, error) {
	if len(zahlen) == 0 {
		return 0, errors.New("Slice ist leer")
	}
	max := zahlen[0]
	for _, z := range zahlen {
		if z > max {
			max = z
		}
	}
	return max, nil
}

// ---------------------------------------------------------------------------
// Übungen
//
// 1. Schreibe Tests für die IstPrimzahl-Funktion (main_test.go)
// 2. Füge einen Test für den Fehlerfall von Durchschnitt hinzu
// 3. Schreibe einen Benchmark für die Fibonacci-Funktion
// 4. Erstelle einen neuen Test für eine selbstgeschriebene Funktion
// 5. Prüfe die Testabdeckung: go test -cover .
// ---------------------------------------------------------------------------
