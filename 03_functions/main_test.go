package main

import "testing"

func TestBegruesse(t *testing.T) {
	// begruessung prints to stdout, so we just verify it doesn't panic
	begruessung("Test")
}

func TestQuadrieren(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{7, 49},
		{-3, 9},
	}
	for _, tt := range tests {
		if got := quadrieren(tt.input); got != tt.want {
			t.Errorf("quadrieren(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestAddieren(t *testing.T) {
	if got := addieren(10, 25); got != 35 {
		t.Errorf("addieren(10, 25) = %d, want 35", got)
	}
	if got := addieren(-5, 5); got != 0 {
		t.Errorf("addieren(-5, 5) = %d, want 0", got)
	}
}

func TestRechteck(t *testing.T) {
	flaeche, umfang := rechteck(5.0, 3.0)
	if flaeche != 15.0 {
		t.Errorf("rechteck Fläche = %f, want 15.0", flaeche)
	}
	if umfang != 16.0 {
		t.Errorf("rechteck Umfang = %f, want 16.0", umfang)
	}
}

func TestTeilen(t *testing.T) {
	quotient, err := teilen(10, 2)
	if err != nil {
		t.Fatalf("unerwarteter Fehler: %v", err)
	}
	if quotient != 5.0 {
		t.Errorf("teilen(10, 2) = %f, want 5.0", quotient)
	}

	_, err = teilen(10, 0)
	if err == nil {
		t.Error("teilen(10, 0) sollte einen Fehler zurückgeben")
	}
}

func TestSummiere(t *testing.T) {
	if got := summiere(1, 2, 3, 4, 5); got != 15 {
		t.Errorf("summiere(1..5) = %d, want 15", got)
	}
	if got := summiere(); got != 0 {
		t.Errorf("summiere() = %d, want 0", got)
	}
}
