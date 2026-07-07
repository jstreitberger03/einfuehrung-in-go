// Kapitel 7: Fehlerbehandlung — Errors, Panic, Defer
//
// 🎯 Lernziele:
// - Das error-Interface verstehen
// - Fehler mit fmt.Errorf erzeugen
// - Sentinelfehler (vordefinierte Fehler)
// - panic und recover für Ausnahmesituationen
// - defer für Aufräumarbeiten
//
// 📖 Los geht's!

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// ==========================================
	// 1. Das error-Interface
	// ==========================================
	fmt.Println("=== Error-Interface ===")

	// error ist ein eingebautes Interface:
	// type error interface {
	//     Error() string
	// }

	datei, err := os.Open("nicht_existent.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen:", err)
	} else {
		datei.Close()
	}

	// ==========================================
	// 2. Eigene Fehler erzeugen
	// ==========================================
	fmt.Println("\n=== Fehler erzeugen ===")

	alter := 15
	err = alterPruefen(alter)
	if err != nil {
		fmt.Println(err)
	}

	_, err = quadratwurzel(-9)
	if err != nil {
		fmt.Println(err)
	}

	// ==========================================
	// 3. Sentinelfehler (vordefinierte Fehler)
	// ==========================================
	fmt.Println("\n=== Sentinelfehler ===")

	_, err = kontostandAbrufen(-1)
	if errors.Is(err, ErrUngueltigeID) {
		fmt.Println("❌ Benutzer-ID ungültig!")
	} else if err != nil {
		fmt.Println("❌ Unbekannter Fehler:", err)
	} else {
		fmt.Println("✅ Kontostand abgerufen")
	}

	_, err = kontostandAbrufen(999)
	if errors.Is(err, ErrBenutzerNichtGefunden) {
		fmt.Println("❌ Benutzer nicht gefunden!")
	} else if err != nil {
		fmt.Println("❌ Unbekannter Fehler:", err)
	}

	// ==========================================
	// 4. defer — Aufräumarbeiten
	// ==========================================
	fmt.Println("\n=== Defer ===")

	dateiBeispiel()
	// defer wird am Ende der Funktion ausgeführt, auch bei panic!
	// Perfekt für: Dateien schließen, Mutex freigeben, Verbindungen trennen

	// ==========================================
	// 5. panic und recover
	// ==========================================
	fmt.Println("\n=== Panic & Recover ===")

	sicherBerechnen(10, 2)
	sicherBerechnen(10, 0) // Kein Absturz!
	sicherBerechnen(8, 4)

	fmt.Println("\n✅ Programm läuft weiter, obwohl eine Division durch 0 versucht wurde!")

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Fehler erzeugen ---

// errors.New erzeugt einen einfachen Fehler
func alterPruefen(alter int) error {
	if alter < 18 {
		return errors.New("Zugriff verweigert: Minderjährig")
	}
	return nil
}

// fmt.Errorf erzeugt einen formatierten Fehler
func quadratwurzel(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("quadratwurzel von %.0f ist nicht definiert", x)
	}
	// Echte Berechnung übersprungen (der Einfachheit halber)
	return 0, nil
}

// --- Sentinelfehler ---

// Vordefinierte Fehler (Konvention: Err-Präfix)
var (
	ErrUngueltigeID         = errors.New("ungültige Benutzer-ID")
	ErrBenutzerNichtGefunden = errors.New("Benutzer nicht gefunden")
	ErrKontoGesperrt        = errors.New("Konto ist gesperrt")
)

type BankKonto struct {
	BenutzerID int
	Kontostand float64
	Gesperrt   bool
}

func kontostandAbrufen(benutzerID int) (float64, error) {
	if benutzerID < 0 {
		return 0, ErrUngueltigeID
	}
	if benutzerID > 100 {
		return 0, ErrBenutzerNichtGefunden
	}
	return 1500.00, nil
}

// --- Defer Beispiel ---

func dateiBeispiel() {
	fmt.Println("Öffne Datei...")
	datei, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	// defer = führe diesen Befehl am Ende der Funktion aus
	defer datei.Close()
	fmt.Println("✅ Datei geöffnet. (Wird automatisch geschlossen)")

	// ... hier könnte noch Code stehen ...
	fmt.Println("Arbeite mit der Datei...")
	// datei.Close() wird automatisch aufgerufen!
}

// --- Panic und Recover ---

func teilenMitPanic(a, b int) int {
	if b == 0 {
		panic("Division durch Null!")
	}
	return a / b
}

func sicherBerechnen(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("⚠️ Panic abgefangen: %v (bei %d/%d)\n", r, a, b)
		}
	}()

	ergebnis := teilenMitPanic(a, b)
	fmt.Printf("%d / %d = %d\n", a, b, ergebnis)
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Schreibe eine Funktion, die einen String in int umwandelt
//    (mit strconv.Atoi) und eigene Fehler kapselt
// 2. Erstelle eine Funktion "kontoBearbeiten", die mehrere Fehlerarten
//    mit errors.Is unterscheiden kann
// 3. Schreibe eine Funktion mit defer, die eine Nachricht "Funktion beendet"
//    ausgibt – egal, ob ein Fehler auftrat
// 4. Erstelle eine Funktion "findPerson", die einen benutzerdefinierten
//    Fehler mit zusätzlichen Informationen zurückgibt (Tipp: struct + error)
// ---------------------------------------------------------------------------
