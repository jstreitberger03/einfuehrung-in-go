// Kapitel 2: Kontrollstrukturen — Entscheidungen treffen
//
// Lernziele:
// - Bedingungen mit if/else
// - Switch-Anweisungen
// - For-Schleifen (die einzige Schleifenart in Go!)
// - Schleifen mit break und continue steuern
//
// Los geht's!

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// ==========================================
	// 1. If / Else
	// ==========================================
	fmt.Println("=== If / Else ===")

	alter := 18

	if alter >= 18 {
		fmt.Println("Du bist volljährig ")
	} else {
		fmt.Println("Du bist noch minderjährig")
	}

	// --- If mit else if ---
	note := 2.0

	if note == 1.0 {
		fmt.Println("Sehr gut!")
	} else if note <= 2.5 {
		fmt.Println("Gut!")
	} else if note <= 3.5 {
		fmt.Println("Befriedigend")
	} else {
		fmt.Println("Nachbesserung empfohlen")
	}

	// --- Kurze If-Anweisung ---
	// Go erlaubt einen kurzen Ausdruck vor der Bedingung:
	if x := math.Sqrt(144); x == 12 {
		fmt.Printf("Die Wurzel von 144 ist %.0f \n", x)
	}
	// x ist außerhalb des if nicht mehr verfügbar!

	// ==========================================
	// 2. Switch
	// ==========================================
	fmt.Println("\n=== Switch ===")

	wochentag := "Samstag"

	switch wochentag {
	case "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag":
		fmt.Println("Arbeitstag ")
	case "Samstag", "Sonntag":
		fmt.Println("Wochenende ")
	default:
		fmt.Println("Ungültiger Wochentag")
	}

	// --- Switch ohne Bedingung (wie if/else) ---
	stunde := time.Now().Hour()

	switch {
	case stunde < 12:
		fmt.Println("Guten Morgen! ")
	case stunde < 17:
		fmt.Println("Guten Tag! ")
	default:
		fmt.Println("Guten Abend! ")
	}

	// ==========================================
	// 3. For-Schleifen
	// ==========================================
	fmt.Println("\n=== For-Schleifen ===")

	// --- Klassische for-Schleife ---
	fmt.Println("Zählen von 1 bis 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// --- For wie while (nur Bedingung) ---
	fmt.Println("\nPotenz von 2 bis 64:")
	zahl := 2
	for zahl <= 64 {
		fmt.Printf("%d ", zahl)
		zahl *= 2
	}
	fmt.Println()

	// --- Endlosschleife (mit break) ---
	fmt.Println("\nEndlosschleife mit break:")
	summe := 0
	for {
		summe += 5
		fmt.Printf("%d ", summe)
		if summe >= 30 {
			break // bricht die Schleife ab
		}
	}
	fmt.Println()

	// --- Range: Über Sammlungen iterieren ---
	fmt.Println("\nRange über eine Liste:")
	fruechte := []string{"Apfel", "Banane", "Kirsche"}
	for i, frucht := range fruechte {
		fmt.Printf("%d: %s\n", i, frucht)
	}

	// Wenn du den Index nicht brauchst, verwende _:
	for _, frucht := range fruechte {
		fmt.Printf("- %s\n", frucht)
	}

	// --- Continue: Überspringe einen Durchlauf ---
	fmt.Println("\nUngerade Zahlen von 1 bis 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // springe zum nächsten Durchlauf
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// ---------------------------------------------------------------------------
// Übungen
//
// 1. Schreibe ein Programm, das prüft, ob eine Zahl gerade oder ungerade ist
// 2. Gib die 7er-Reihe aus: 7, 14, 21, ..., 70
// 3. Schreibe eine Schleife, die von 10 bis 1 rückwärts zählt
// 4. Erstelle einen Switch für eine Note (1-5) und gib die Bewertung aus
// 5. Finde alle Zahlen von 1-100, die durch 3 UND 5 teilbar sind (FizzBuzz)
// ---------------------------------------------------------------------------
