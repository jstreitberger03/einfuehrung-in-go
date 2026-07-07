// Paket taschenrechner – einfache mathematische Operationen
package taschenrechner

import "math"

// Exportierte Namen beginnen mit Großbuchstaben
// Nicht-exportierte (private) Namen mit Kleinbuchstaben

// Pi ist ein exportierter Wert (≈ 3.14159)
const Pi = 3.141592653589793

// pi ist ein nicht-exportierter Wert (nur innerhalb des Pakets sichtbar)
const pi = 3.14

// Addiere berechnet a + b
func Addiere(a, b int) int {
	return a + b
}

// Subtrahiere berechnet a - b
func Subtrahiere(a, b int) int {
	return a - b
}

// verdoppeln ist nicht-exportiert (beginnt mit Kleinbuchstaben)
func verdoppeln(x int) int {
	return x * 2
}

// Wurzel berechnet die Quadratwurzel
func Wurzel(x float64) float64 {
	return math.Sqrt(x)
}

// Potenz berechnet x^y
func Potenz(x, y float64) float64 {
	return math.Pow(x, y)
}
