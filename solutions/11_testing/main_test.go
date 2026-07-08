package main

import "testing"

// Übung 1: Tests für IstPrimzahl
func TestIstPrimzahl(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"0 ist nicht prim", 0, false},
		{"1 ist nicht prim", 1, false},
		{"2 ist prim", 2, true},
		{"3 ist prim", 3, true},
		{"4 ist nicht prim", 4, false},
		{"7 ist prim", 7, true},
		{"9 ist nicht prim", 9, false},
		{"13 ist prim", 13, true},
		{"negativ nicht prim", -5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IstPrimzahl(tt.n); got != tt.want {
				t.Errorf("IstPrimzahl(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// Übung 2: Test für Fehlerfall von Durchschnitt
func TestDurchschnitt_Leer(t *testing.T) {
	_, err := Durchschnitt([]float64{})
	if err == nil {
		t.Error("Durchschnitt([]) sollte einen Fehler zurückgeben")
	}
}

func TestDurchschnitt_OK(t *testing.T) {
	got, err := Durchschnitt([]float64{2, 4, 6})
	if err != nil {
		t.Fatalf("unerwarteter Fehler: %v", err)
	}
	if got != 4.0 {
		t.Errorf("Durchschnitt([2,4,6]) = %f, want 4.0", got)
	}
}

// Übung 3: Benchmark für Fibonacci
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fibonacci(20)
	}
}

// Übung 4: Test für MeineFunktion
func TestMeineFunktion(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{5, true},
		{-1, false},
		{0, false},
	}
	for _, tt := range tests {
		if got := MeineFunktion(tt.n); got != tt.want {
			t.Errorf("MeineFunktion(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// Übung 5: go test -cover .
func TestMax_Leer(t *testing.T) {
	_, err := Max([]int{})
	if err == nil {
		t.Error("Max([]) sollte einen Fehler zurückgeben")
	}
}

func TestMax_OK(t *testing.T) {
	got, err := Max([]int{3, 7, 2, 9, 1})
	if err != nil {
		t.Fatalf("unerwarteter Fehler: %v", err)
	}
	if got != 9 {
		t.Errorf("Max([3,7,2,9,1]) = %d, want 9", got)
	}
}
