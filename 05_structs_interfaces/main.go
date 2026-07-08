// Kapitel 5: Structs und Interfaces — Eigene Typen definieren
//
// Lernziele:
// - Eigene Datentypen mit struct definieren
// - Methoden (Funktionen mit Empfänger)
// - Interfaces (implizite Implementierung)
// - Type Embedding ("Vererbung" in Go)
//
// Los geht's!

package main

import "fmt"

// ==========================================
// 1. Structs
// ==========================================

// Person ist ein eigener Datentyp mit mehreren Feldern
type Person struct {
	Name  string
	Alter int
	Stadt string
	Beruf string
}

func main() {
	// --- Struct-Instanzen erstellen ---
	fmt.Println("=== Structs ===")

	// Variante 1: Mit Feldnamen (empfohlen)
	anna := Person{
		Name:  "Anna",
		Alter: 28,
		Stadt: "Wien",
		Beruf: "Ingenieurin",
	}

	// Variante 2: Ohne Feldnamen (positional)
	ben := Person{"Ben", 34, "München", "Arzt"}

	fmt.Println("Anna:", anna)
	fmt.Println("Ben:", ben)

	// --- Auf Felder zugreifen ---
	fmt.Println("\nAnnas Details:")
	fmt.Println("Name:", anna.Name)
	fmt.Println("Alter:", anna.Alter)
	fmt.Println("Stadt:", anna.Stadt)

	// --- Werte ändern ---
	anna.Alter = 29
	anna.Stadt = "Salzburg"
	fmt.Println("\nNach Umzug:", anna)

	// --- Struct mit匿名 struct ---
	// Für einmalige Verwendung
	adresse := struct {
		Strasse string
		Plz     int
	}{
		Strasse: "Hauptstr. 12",
		Plz:     1010,
	}
	fmt.Println("\nAdresse:", adresse)

	// ==========================================
	// 2. Methoden (Funktionen mit Empfänger)
	// ==========================================
	fmt.Println("\n=== Methoden ===")

	// Methodenaufruf mit Punktnotation
	fmt.Println(anna.vorstellen())
	fmt.Println(ben.vorstellen())

	fmt.Println(anna.geburtstagFeiern()) // Alter +1
	fmt.Println("Neues Alter:", anna.Alter)

	// ==========================================
	// 3. Interfaces (implizit)
	// ==========================================
	fmt.Println("\n=== Interfaces ===")

	// Ein Interface wird automatisch implementiert –
	// kein explizites "implements" nötig!
	tiere := []Lebewesen{
		Hund{Name: "Rex"},
		Katze{Name: "Miau"},
		Mensch{Name: "David"},
	}

	for _, tier := range tiere {
		fmt.Println(tier.geraeuschMachen())
	}

	// ==========================================
	// 4. Type Embedding
	// ==========================================
	fmt.Println("\n=== Type Embedding ===")

	mitarbeiter := Mitarbeiter{
		Person: Person{Name: "Eva", Alter: 32, Stadt: "Berlin", Beruf: "Entwicklerin"},
		Firma:  "TechCorp",
		Gehalt: 75000,
	}
	fmt.Println(mitarbeiter.vorstellen())
	fmt.Printf("Firma: %s, Gehalt: %d €\n", mitarbeiter.Firma, mitarbeiter.Gehalt)

	fmt.Println("\n Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Struct-Definitionen ---

// Vorstellen ist eine Methode auf Person
// Der Empfänger (p Person) steht zwischen func und Methodenname
func (p Person) vorstellen() string {
	return fmt.Sprintf("Hallo, ich bin %s (%d) aus %s und arbeite als %s.",
		p.Name, p.Alter, p.Stadt, p.Beruf)
}

// Pointer-Empfänger: Änderungen sind auch außerhalb sichtbar!
func (p *Person) geburtstagFeiern() string {
	p.Alter++
	return fmt.Sprintf("Happy Birthday, %s! Du bist jetzt %d! ", p.Name, p.Alter)
}

// --- Interfaces ---

// Lebewesen ist ein Interface – ein Satz von Methoden-Signaturen
type Lebewesen interface {
	geraeuschMachen() string
}

// Hund implementiert Lebewesen automatisch
type Hund struct {
	Name string
}

func (h Hund) geraeuschMachen() string {
	return fmt.Sprintf("%s macht: Wuff Wuff! ", h.Name)
}

// Katze implementiert Lebewesen automatisch
type Katze struct {
	Name string
}

func (k Katze) geraeuschMachen() string {
	return fmt.Sprintf("%s macht: Miau Miau! ", k.Name)
}

// Mensch implementiert Lebewesen automatisch
type Mensch struct {
	Name string
}

func (m Mensch) geraeuschMachen() string {
	return fmt.Sprintf("%s sagt: Hallo! ", m.Name)
}

// --- Type Embedding ---

// Mitarbeiter embeded Person – alle Person-Methoden sind verfügbar!
type Mitarbeiter struct {
	Person // Embedding: Person wird eingebettet
	Firma  string
	Gehalt int
}

// Überschreibt die vorstellen()-Methode von Person
func (m Mitarbeiter) vorstellen() string {
	return fmt.Sprintf("%s arbeitet bei %s und verdient %d €.",
		m.Name, m.Firma, m.Gehalt)
}

// ---------------------------------------------------------------------------
// Übungen
//
// 1. Definiere ein struct "Buch" mit den Feldern: Titel, Autor, Seiten, Jahr
// 2. Erstelle eine Methode "info()" für Buch, die eine Zusammenfassung zurückgibt
// 3. Definiere ein Interface "Fahrzeug" mit Methode "bewegen() string"
// 4. Implementiere Fahrzeug mit den Typen Auto und Fahrrad
// 5. Erstelle ein struct "EBook", das Buch embeded und ein Feld "DateigroesseMB" hat
// ---------------------------------------------------------------------------
