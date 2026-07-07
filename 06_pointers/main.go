// Kapitel 6: Pointer — Referenzen verstehen
//
// 🎯 Lernziele:
// - Was sind Pointer? Adressen im Speicher
// - Pointer deklarieren und dereferenzieren (*, &)
// - nil vermeiden und prüfen
// - Pointer als Funktionsparameter
// - Stack vs. Heap (kurzer Überblick)
//
// 📖 Los geht's!

package main

import "fmt"

func main() {
	// ==========================================
	// 1. Grundlagen: & und *
	// ==========================================
	fmt.Println("=== Pointer Grundlagen ===")

	x := 42
	p := &x // & = Adresse von x (Pointer auf x)

	fmt.Println("Wert von x:", x)
	fmt.Println("Adresse von x:", &x)
	fmt.Println("Pointer p:", p)      // speichert die Adresse von x
	fmt.Println("Wert via *p:", *p)   // * = Dereferenzierung: Wert an der Adresse

	// --- Wert über Pointer ändern ---
	*p = 100                         // Ändere den Wert an der Adresse von x
	fmt.Println("\nNach *p = 100:")
	fmt.Println("x =", x)            // 100 – x wurde geändert!
	fmt.Println("*p =", *p)          // 100

	// ==========================================
	// 2. Pointer als Funktionsparameter
	// ==========================================
	fmt.Println("\n=== Pointer als Parameter ===")

	zahl := 10
	fmt.Println("Vor verdoppeln:", zahl)

	// Übergabe als Wert (Kopie) – nichts passiert
	verdoppleWert(zahl)
	fmt.Println("Nach verdoppleWert:", zahl) // 10 (unverändert!)

	// Übergabe als Pointer
	verdopplePointer(&zahl)
	fmt.Println("Nach verdopplePointer:", zahl) // 20

	// ==========================================
	// 3. nil – der Null-Pointer
	// ==========================================
	fmt.Println("\n=== nil ===")

	var ptr *int // nil (noch keine Adresse)
	fmt.Println("ptr ist nil:", ptr == nil)

	// Achtung: Dereferenzierung von nil = panic!
	// fmt.Println(*ptr) // 💥 panic: runtime error

	// Sicherer Zugriff:
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Pointer ist nil – kein Zugriff möglich")
	}

	// ==========================================
	// 4. nil und Funktionen
	// ==========================================
	ergebnis := safeDivide(10, 2)
	if ergebnis != nil {
		fmt.Println("\n10 / 2 =", *ergebnis)
	}

	ergebnis2 := safeDivide(10, 0)
	if ergebnis2 != nil {
		fmt.Println("10 / 0 =", *ergebnis2)
	} else {
		fmt.Println("10 / 0 = nicht definiert (nil zurückgegeben)")
	}

	// ==========================================
	// 5. Pointer vs. Value Receiver
	// ==========================================
	fmt.Println("\n=== Pointer vs. Value Receiver ===")

	kreis := Kreis{Radius: 5.0}
	fmt.Println("Kreis:", kreis)

	kreis.vergroessern(2.0)
	fmt.Println("Nach vergrössern:", kreis)

	// ==========================================
	// 6. Wann Pointer verwenden? (Faustregeln)
	// ==========================================
	fmt.Println("\n=== Faustregeln für Pointer ===")
	fmt.Println("✅ Pointer verwenden, wenn:")
	fmt.Println("  - Du Werte in einer Funktion ändern willst")
	fmt.Println("  - Große Structs vermeiden kopiert zu werden")
	fmt.Println("  - nil als Zustand erlaubt sein soll")
	fmt.Println()
	fmt.Println("❌ Kein Pointer, wenn:")
	fmt.Println("  - Kleine, unveränderliche Werte (int, bool, float64)")
	fmt.Println("  - nil keinen Sinn ergibt")
	fmt.Println("  - Du den Wert nicht ändern musst")

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Hilfsfunktionen ---

// Wert-Übergabe: zahl ist eine Kopie!
func verdoppleWert(zahl int) {
	zahl *= 2
	fmt.Println("Innerhalb:", zahl) // verdoppelt
}

// Pointer-Übergabe: zahl ist der originale Wert
func verdopplePointer(zahl *int) {
	*zahl *= 2 // Ändert den Original-Wert
}

// Pointer als Rückgabewert (nil für Fehlerfall)
func safeDivide(a, b float64) *float64 {
	if b == 0 {
		return nil // nil bedeutet "kein gültiges Ergebnis"
	}
	ergebnis := a / b
	return &ergebnis
}

// --- Beispiel mit Struct ---

type Kreis struct {
	Radius float64
}

// Value Receiver: kreis ist eine Kopie
func (k Kreis) flaeche() float64 {
	return 3.14159 * k.Radius * k.Radius
}

// Pointer Receiver: Änderungen wirken sich aus
func (k *Kreis) vergroessern(faktor float64) {
	k.Radius *= faktor
}

func (k Kreis) String() string {
	return fmt.Sprintf("Kreis (r=%.1f, fläche=%.2f)", k.Radius, k.flaeche())
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Schreibe eine Funktion "swap(a, b *int)", die zwei Werte vertauscht
// 2. Schreibe eine Funktion, die einen *string akzeptiert und den Wert
//    in Großbuchstaben umwandelt (Hinweis: strings.ToUpper)
// 3. Erstelle ein Struct "Bankkonto" mit Methode "einzahlen(amount *float64)"
// 4. Warum gibt es in Go keine call-by-reference? (Denkaufgabe)
// ---------------------------------------------------------------------------
