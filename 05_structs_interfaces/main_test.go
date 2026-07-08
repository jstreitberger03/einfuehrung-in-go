package main

import (
	"strings"
	"testing"
)

func TestPersonVorstellen(t *testing.T) {
	p := Person{Name: "Test", Alter: 30, Stadt: "Berlin", Beruf: "Dev"}
	got := p.vorstellen()
	if !strings.Contains(got, "Test") {
		t.Errorf("vorstellen() enthält nicht den Namen: %s", got)
	}
	if !strings.Contains(got, "Berlin") {
		t.Errorf("vorstellen() enthält nicht die Stadt: %s", got)
	}
}

func TestGeburtstagFeiern(t *testing.T) {
	p := Person{Name: "Anna", Alter: 28}
	msg := p.geburtstagFeiern()
	if p.Alter != 29 {
		t.Errorf("Alter nach Geburtstag = %d, want 29", p.Alter)
	}
	if !strings.Contains(msg, "Anna") {
		t.Errorf("Geburtstagsnachricht enthält nicht den Namen: %s", msg)
	}
}

func TestLebewesenInterface(t *testing.T) {
	tiere := []Lebewesen{
		Hund{Name: "Rex"},
		Katze{Name: "Miau"},
		Mensch{Name: "David"},
	}
	for _, tier := range tiere {
		msg := tier.geraeuschMachen()
		if msg == "" {
			t.Error("geraeuschMachen() darf nicht leer sein")
		}
	}
}

func TestMitarbeiterEmbedding(t *testing.T) {
	m := Mitarbeiter{
		Person: Person{Name: "Eva", Alter: 32, Stadt: "Berlin", Beruf: "Dev"},
		Firma:  "TechCorp",
		Gehalt: 75000,
	}
	// Embedding: Zugriff auf Person-Felder
	if m.Name != "Eva" {
		t.Errorf("m.Name = %q, want %q", m.Name, "Eva")
	}
	// Überschriebene Methode
	msg := m.vorstellen()
	if !strings.Contains(msg, "TechCorp") {
		t.Errorf("vorstellen() enthält nicht die Firma: %s", msg)
	}
}
