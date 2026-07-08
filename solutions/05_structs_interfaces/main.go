package main

import "fmt"

// Übung 1: Struct Buch
type Buch struct {
	Titel  string
	Autor  string
	Seiten int
	Jahr   int
}

// Übung 2: Methode info() für Buch
func (b Buch) info() string {
	return fmt.Sprintf("\"%s\" von %s (%d), %d Seiten", b.Titel, b.Autor, b.Jahr, b.Seiten)
}

// Übung 3: Interface Fahrzeug
type Fahrzeug interface {
	bewegen() string
}

// Übung 4: Auto und Fahrrad implementieren Fahrzeug
type Auto struct {
	Marke string
}

func (a Auto) bewegen() string {
	return fmt.Sprintf("%s fährt auf der Straße", a.Marke)
}

type Fahrrad struct {
	Typ string
}

func (f Fahrrad) bewegen() string {
	return fmt.Sprintf("%s radelt auf dem Radweg", f.Typ)
}

// Übung 5: EBook embeddet Buch
type EBook struct {
	Buch
	DateigroesseMB float64
}

func main() {
	// Übung 1+2: Buch mit Methode
	b := Buch{Titel: "Go Programming", Autor: "Alan Donovan", Seiten: 380, Jahr: 2015}
	fmt.Println(b.info())

	// Übung 3+4: Interface Fahrzeug
	fahrzeuge := []Fahrzeug{
		Auto{Marke: "BMW"},
		Fahrrad{Typ: "Mountainbike"},
	}
	for _, f := range fahrzeuge {
		fmt.Println(f.bewegen())
	}

	// Übung 5: EBook mit Embedding
	ebook := EBook{
		Buch:           Buch{Titel: "Go in Action", Autor: "William Kennedy", Seiten: 280, Jahr: 2016},
		DateigroesseMB: 2.5,
	}
	fmt.Printf("\n%s (%.1f MB)\n", ebook.info(), ebook.DateigroesseMB)
}
