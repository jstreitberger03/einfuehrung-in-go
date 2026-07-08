package main

import (
	"errors"
	"fmt"
)

func main() {
	// Übung 1: isEven
	fmt.Println("IsEven(4):", isEven(4))
	fmt.Println("IsEven(7):", isEven(7))

	// Übung 2: Celsius zu Fahrenheit
	fmt.Printf("\n20°C = %.1f°F\n", celsiusToFahrenheit(20))
	fmt.Printf("0°C = %.1f°F\n", celsiusToFahrenheit(0))
	fmt.Printf("100°C = %.1f°F\n", celsiusToFahrenheit(100))

	// Übung 3: minMax
	min, max := minMax(3, 7, 5)
	fmt.Printf("\nminMax(3, 7, 5): min=%d, max=%d\n", min, max)

	// Übung 4: Durchschnitt (variadic)
	durchschnitt := average(10, 20, 30, 40, 50)
	fmt.Printf("\nDurchschnitt von [10,20,30,40,50]: %.1f\n", durchschnitt)

	// Übung 5: Fakultät mit Fehlerbehandlung
	for _, n := range []int{0, 1, 5, -3} {
		ergebnis, err := factorial(n)
		if err != nil {
			fmt.Printf("factorial(%d): Fehler: %v\n", n, err)
		} else {
			fmt.Printf("factorial(%d) = %d\n", n, ergebnis)
		}
	}
}

// Übung 1: Prüft, ob eine Zahl gerade ist
func isEven(n int) bool {
	return n%2 == 0
}

// Übung 2: Celsius in Fahrenheit umrechnen
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// Übung 3: Minimum und Maximum von drei Zahlen
func minMax(a, b, c int) (min, max int) {
	min = a
	max = a
	if b < min {
		min = b
	}
	if b > max {
		max = b
	}
	if c < min {
		min = c
	}
	if c > max {
		max = c
	}
	return
}

// Übung 4: Variadic Durchschnittsberechnung
func average(zahlen ...float64) float64 {
	if len(zahlen) == 0 {
		return 0
	}
	summe := 0.0
	for _, z := range zahlen {
		summe += z
	}
	return summe / float64(len(zahlen))
}

// Übung 5: Fakultät mit Fehlerbehandlung
func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Fakultät für negative Zahlen nicht definiert")
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
