// Kapitel 8: Packages & Module — Code organisieren
//
// 🎯 Lernziele:
// - Eigene Pakete erstellen und importieren
// - Exportierte vs. nicht-exportierte Namen
// - go.mod verstehen
// - Externe Pakete installieren
//
// 📖 Los geht's!

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/julianstreitberger/einfuehrung-in-go/08_packages_modules/taschenrechner"
	"github.com/julianstreitberger/einfuehrung-in-go/08_packages_modules/textutils"
)

func main() {
	// ==========================================
	// 1. Eigene Pakete verwenden
	// ==========================================
	fmt.Println("=== Eigene Pakete ===")

	// Funktion aus dem taschenrechner-Paket
	summe := taschenrechner.Addiere(10, 20)
	fmt.Println("10 + 20 =", summe)

	differenz := taschenrechner.Subtrahiere(50, 13)
	fmt.Println("50 - 13 =", differenz)

	// Funktion aus dem textutils-Paket
	begruessung := textutils.Begruesse("Anna")
	fmt.Println(begruessung)

	text := "hallo go welt!"
	gross := textutils.Grossschreiben(text)
	fmt.Println("Original:", text)
	fmt.Println("Groß:    ", gross)

	// ==========================================
	// 2. Exportierte vs. nicht-exportierte Namen
	// ==========================================
	fmt.Println("\n=== Exportierte vs. Nicht-Exportierte Namen ===")

	// ✅ Exportiert: beginnt mit Großbuchstaben
	fmt.Println("✅ taschenrechner.Pi =", taschenrechner.Pi)

	// ❌ Nicht exportiert: beginnt mit Kleinbuchstaben
	// fmt.Println(taschenrechner.pi) // Compiler-Fehler!
	// fmt.Println(taschenrechner.verdoppeln(5)) // Compiler-Fehler!

	// ==========================================
	// 3. Die Standardbibliothek
	// ==========================================
	fmt.Println("\n=== Standardbibliothek ===")

	// Go hat eine umfangreiche Standardbibliothek:
	// - fmt, strings, math, os, time, json, http ...
	fmt.Println("math.Sqrt(144) =", math.Sqrt(144))
	fmt.Println("strings.ToUpper('go') =", strings.ToUpper("go"))
	fmt.Println("strings.Repeat('Go! ', 3) =", strings.Repeat("Go! ", 3))

	// ==========================================
	// 4. go.mod verstehen
	// ==========================================
	fmt.Println("\n=== go.mod ===")
	fmt.Println("Die Datei go.mod definiert:")
	fmt.Println("- Den Modul-Pfad (Namespace)")
	fmt.Println("- Die Go-Version")
	fmt.Println("- Externe Abhängigkeiten")
	fmt.Println()
	fmt.Println("Wichtige Befehle:")
	fmt.Println("  go mod init <name>     — Modul initialisieren")
	fmt.Println("  go mod tidy            — Abhängigkeiten bereinigen")
	fmt.Println("  go get <paket>         — Paket installieren")
	fmt.Println("  go list -m all         — Alle Abhängigkeiten anzeigen")

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Erstelle ein neues Paket "wetter" (in einem Unterordner) mit einer
//    Funktion "TemperaturBeschreibung(temp float64) string"
// 2. Füge eine nicht-exportierte Hilfsfunktion in taschenrechner hinzu
// 3. Installiere ein externes Paket mit "go get" und verwende es
//    (z.B. github.com/fatih/color für farbige Ausgabe)
// 4. Erstelle eine init()-Funktion in einem Paket und beobachte, wann
//    sie automatisch aufgerufen wird
// ---------------------------------------------------------------------------
