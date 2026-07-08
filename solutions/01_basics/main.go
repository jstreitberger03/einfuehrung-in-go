package main

import "fmt"

func main() {
	// Übung 1: Variablen deklarieren
	lieblingszahl := 42
	lieblingsfarbe := "Blau"
	geburtsjahr := 1999

	fmt.Println("Lieblingszahl:", lieblingszahl)
	fmt.Println("Lieblingsfarbe:", lieblingsfarbe)
	fmt.Println("Geburtsjahr:", geburtsjahr)

	// Übung 2: Alter berechnen
	alter := 2026 - geburtsjahr
	fmt.Println("Alter (ca.):", alter)

	// Übung 3: Konstante Lichtgeschwindigkeit
	const lichtgeschwindigkeit = 299_792_458 // m/s
	fmt.Printf("Lichtgeschwindigkeit: %d m/s\n", lichtgeschwindigkeit)

	// Übung 4: fmt.Printf mit Name und Alter
	name := "Julian"
	fmt.Printf("Ich heiße %s und bin %d Jahre alt.\n", name, alter)

	// Übung 5: Zwei Variablen tauschen
	a := 5
	b := 10
	fmt.Printf("Vor dem Tausch: a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("Nach dem Tausch: a=%d, b=%d\n", a, b)
}
