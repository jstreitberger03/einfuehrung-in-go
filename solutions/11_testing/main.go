package main

import (
	"errors"
	"fmt"
	"math"
)

// MeineFunktion prüft, ob eine Zahl positiv ist (Übung 4)
func MeineFunktion(n int) bool {
	return n > 0
}

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

func main() {
	fmt.Println("Die Lösungen zu Kapitel 11 findest du in main_test.go.")
	fmt.Println()
	fmt.Println("Übung 1: Tests für IstPrimzahl (siehe main_test.go)")
	fmt.Println("Übung 2: Test für Fehlerfall von Durchschnitt")
	fmt.Println("Übung 3: Benchmark für Fibonacci")
	fmt.Println("Übung 4: Test für MeineFunktion")
	fmt.Println("Übung 5: go test -cover .")
}
