package main

import (
	"errors"
	"testing"
)

func TestAlterPruefen(t *testing.T) {
	if err := alterPruefen(25); err != nil {
		t.Errorf("alterPruefen(25) sollte nil sein, got %v", err)
	}
	if err := alterPruefen(15); err == nil {
		t.Error("alterPruefen(15) sollte einen Fehler zurückgeben")
	}
}

func TestQuadratwurzel(t *testing.T) {
	_, err := quadratwurzel(-9)
	if err == nil {
		t.Error("quadratwurzel(-9) sollte einen Fehler zurückgeben")
	}
}

func TestKontostandAbrufen(t *testing.T) {
	_, err := kontostandAbrufen(-1)
	if !errors.Is(err, ErrUngueltigeID) {
		t.Errorf("erwartet ErrUngueltigeID, got %v", err)
	}

	_, err = kontostandAbrufen(999)
	if !errors.Is(err, ErrBenutzerNichtGefunden) {
		t.Errorf("erwartet ErrBenutzerNichtGefunden, got %v", err)
	}

	betrag, err := kontostandAbrufen(50)
	if err != nil {
		t.Fatalf("unerwarteter Fehler: %v", err)
	}
	if betrag != 1500.00 {
		t.Errorf("kontostandAbrufen(50) = %f, want 1500.00", betrag)
	}
}

func TestSentinelErrors(t *testing.T) {
	err := ErrUngueltigeID
	if err.Error() == "" {
		t.Error("ErrUngueltigeID.Error() darf nicht leer sein")
	}
	if ErrKontoGesperrt == nil {
		t.Error("ErrKontoGesperrt darf nicht nil sein")
	}
}
