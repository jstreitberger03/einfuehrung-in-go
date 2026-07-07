// Kapitel 3: Funktionen — Wiederverwendbarer Code
//
// 🎯 Lernziele:
// - Funktionen definieren und aufrufen
// - Parameter und Rückgabewerte
// - Mehrfachrückgaben (Go-Spezialität!)
// - Benannte Rückgabewerte
// - Variadic Functions (variable Parameteranzahl)
//
// 📖 Los geht's!

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Funktionen ===")

	// --- Einfacher Funktionsaufruf ---
	begruessung("Anna")      // Hallo, Anna!
	begruessung("Ben")       // Hallo, Ben!
	begruessung("Clara")     // Hallo, Clara!

	// --- Funktion mit Rückgabewert ---
	quadrat := quadrieren(7)
	fmt.Println("\n7² =", quadrat)

	// --- Funktion mit mehreren Parametern ---
	ergebnis := addieren(10, 25)
	fmt.Println("10 + 25 =", ergebnis)

	// --- Mehrfachrückgaben ---
	flaeche, umfang := rechteck(5.0, 3.0)
	fmt.Printf("\nRechteck (5x3): Fläche=%.1f, Umfang=%.1f\n", flaeche, umfang)

	// --- Fehler als Rückgabewert ---
	quotient, err := teilen(10, 2)
	if err != nil {
		fmt.Println("Fehler:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	quotient2, err2 := teilen(10, 0)
	if err2 != nil {
		fmt.Println("Fehler:", err2)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", quotient2)
	}

	// --- Blank Identifier (_) ---
	// Wenn du einen Rückgabewert ignorieren willst:
	flaeche2, _ := rechteck(4.0, 2.5)
	fmt.Printf("\nFläche (4x2.5) = %.1f (Umfang ignoriert)\n", flaeche2)

	// --- Benannte Rückgabewerte ---
	x, y := kreis(3.0)
	fmt.Printf("Kreis (r=3): Umfang=%.2f, Fläche=%.2f\n", x, y)

	// --- Variadic Functions ---
	summe := summiere(1, 2, 3, 4, 5)
	fmt.Println("\nSumme von 1-5:", summe)

	// Ein Slice kann mit ... entpackt werden:
	zahlen := []int{10, 20, 30}
	summe2 := summiere(zahlen...)
	fmt.Println("Summe von [10, 20, 30]:", summe2)

	// --- Funktionen als Werte ---
	// In Go sind Funktionen "first-class citizens":
	verdoppler := func(x int) int {
		return x * 2
	}
	fmt.Println("\nVerdoppler (21):", verdoppler(21))
}

// --- Funktionsdefinitionen ---

// Einfache Funktion ohne Rückgabewert
func begruessung(name string) {
	fmt.Printf("Hallo, %s!\n", name)
}

// Funktion mit einem Parameter und einem Rückgabewert
func quadrieren(x int) int {
	return x * x
}

// Funktion mit zwei Parametern
func addieren(a, b int) int {
	// Wenn Parameter denselben Typ haben:
	// func addieren(a, b int) statt func addieren(a int, b int)
	return a + b
}

// Mehrfachrückgabe: (float64, float64)
func rechteck(laenge, breite float64) (float64, float64) {
	flaeche := laenge * breite
	umfang := 2 * (laenge + breite)
	return flaeche, umfang
}

// Funktion mit Fehler-Rückgabe
func teilen(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division durch Null nicht erlaubt")
	}
	return a / b, nil
}

// Benannte Rückgabewerte
func kreis(radius float64) (umfang, flaeche float64) {
	// Die Rückgabevariablen sind bereits deklariert!
	umfang = 2 * math.Pi * radius
	flaeche = math.Pi * radius * radius
	return // "nacktes" return — gibt umfang und flaeche zurück
}

// Variadic Function: ...int = beliebig viele int-Parameter
func summiere(zahlen ...int) int {
	summe := 0
	for _, zahl := range zahlen {
		summe += zahl
	}
	return summe
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Schreibe eine Funktion isEven(n int) bool, die prüft, ob n gerade ist
// 2. Schreibe eine Funktion celsiusToFahrenheit(c float64) float64
// 3. Schreibe eine Funktion minMax(a, b, c int) (min, max int)
// 4. Schreibe eine variadic Funktion, die den Durchschnitt berechnet
// 5. Schreibe eine Funktion factorial(n int) (int, error), die die Fakultät berechnet
//    (Fehler bei negativen Zahlen)
// ---------------------------------------------------------------------------
