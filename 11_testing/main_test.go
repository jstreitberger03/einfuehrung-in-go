// Kapitel 11: Tests — main_test.go
//
// Dieses Datei enthält Tests für die Funktionen in main.go.
// Führe sie aus mit: go test -v .

package main

import "testing"

// ==========================================
// 1. Einfache Unit-Tests
// ==========================================

func TestSumme(t *testing.T) {
	ergebnis := Summe([]int{1, 2, 3, 4, 5})
	erwartet := 15

	if ergebnis != erwartet {
		t.Errorf("Summe([1,2,3,4,5]) = %d; erwartet %d", ergebnis, erwartet)
	}
}

func TestSummeLeer(t *testing.T) {
	ergebnis := Summe([]int{})
	erwartet := 0

	if ergebnis != erwartet {
		t.Errorf("Summe([]) = %d; erwartet %d", ergebnis, erwartet)
	}
}

// ==========================================
// 2. Table-Driven Tests
// ==========================================
//
// Das ist das idiomatische Testmuster in Go!
// Testfälle werden als Liste (Tabelle) definiert.

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		erwartet int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=5", 5, 5},
		{"n=10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ergebnis, err := Fibonacci(tt.n)
			if err != nil {
				t.Errorf("Fibonacci(%d) unerwarteter Fehler: %v", tt.n, err)
			}
			if ergebnis != tt.erwartet {
				t.Errorf("Fibonacci(%d) = %d; erwartet %d", tt.n, ergebnis, tt.erwartet)
			}
		})
	}
}

func TestFibonacciFehler(t *testing.T) {
	_, err := Fibonacci(-1)
	if err == nil {
		t.Error("Fibonacci(-1) sollte einen Fehler zurückgeben")
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		zahlen   []int
		erwartet int
	}{
		{"positive Zahlen", []int{1, 5, 3, 9, 2}, 9},
		{"negative Zahlen", []int{-5, -2, -8, -1}, -1},
		{"ein Element", []int{42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ergebnis, err := Max(tt.zahlen)
			if err != nil {
				t.Errorf("Max(%v) unerwarteter Fehler: %v", tt.zahlen, err)
			}
			if ergebnis != tt.erwartet {
				t.Errorf("Max(%v) = %d; erwartet %d", tt.zahlen, ergebnis, tt.erwartet)
			}
		})
	}
}

// ==========================================
// 3. Benchmarks
// ==========================================

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkSumme(b *testing.B) {
	zahlen := make([]int, 1000)
	for i := range zahlen {
		zahlen[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Summe(zahlen)
	}
}
