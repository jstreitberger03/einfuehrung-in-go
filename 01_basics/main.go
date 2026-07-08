// Kapitel 1: Grundlagen — Variablen, Datentypen, Konstanten
//
// Lernziele:
// - Variablen deklarieren und initialisieren
// - Die wichtigsten Datentypen kennen (int, float64, string, bool)
// - Konstanten verwenden
// - Ein- und Ausgabe mit fmt
// - Typ-Inferenz verstehen (:=)
//
// Los geht's!

package main

import "fmt"

func main() {
	// --- Variablen deklarieren ---
	// var name typ = wert
	var name string = "Anna"
	var alter int = 28
	var groesse float64 = 1.72
	var verheiratet bool = false

	fmt.Println("=== Variablen ===")
	fmt.Println("Name:", name)
	fmt.Println("Alter:", alter)
	fmt.Println("Größe:", groesse)
	fmt.Println("Verheiratet:", verheiratet)

	// --- Kurzschreibweise mit := ---
	// Go kann den Typ aus dem Wert ableiten (Typ-Inferenz)
	stadt := "Wien"
	wohnungsgroesse := 65.5
	kinder := 0

	fmt.Println("\n=== Kurzschreibweise ===")
	fmt.Println("Stadt:", stadt)
	fmt.Println("Wohnungsgröße:", wohnungsgroesse, "m²")
	fmt.Println("Kinder:", kinder)

	// --- Warum := statt var? ---
	// var: wenn du den Typ explizit angeben willst
	// :=: kurze Schreibweise innerhalb von Funktionen
	//
	// Wichtig: := deklariert eine NEUE Variable.
	// = weist einer EXISTIERENDEN Variable einen Wert zu.

	// --- Datentypen im Detail ---
	fmt.Println("\n=== Datentypen ===")

	// Ganze Zahlen (int)
	var einwohner int = 1_950_000 // _ als Tausendertrenner erlaubt!
	fmt.Printf("Einwohner: %d\n", einwohner)

	// Dezimalzahlen (float64)
	var pi float64 = 3.1415926535
	fmt.Printf("Pi: %.4f\n", pi) // %.4f = 4 Nachkommastellen

	// Text (string)
	var begruessung string = "Hallo, Go!"
	fmt.Println("Begrüßung:", begruessung)

	// Wahrheitswert (bool)
	var istCool bool = true
	fmt.Println("Ist Go cool?", istCool)

	// --- Nullwerte (Zero Values) ---
	// In Go hat jede Variable einen Startwert, auch ohne Zuweisung:
	var x int     // 0
	var y float64 // 0
	var z string  // "" (leerer String)
	var w bool    // false

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("int: %d, float64: %.1f, string: %q, bool: %v\n", x, y, z, w)

	// --- Konstanten ---
	// Konstanten können nicht geändert werden!
	const mehrwertsteuer = 0.19
	const piExakt = 3.141592653589793

	fmt.Println("\n=== Konstanten ===")
	preis := 100.0
	steuer := preis * mehrwertsteuer
	fmt.Printf("Preis: %.2f €\n", preis)
	fmt.Printf("Mehrwertsteuer (%.0f%%): %.2f €\n", mehrwertsteuer*100, steuer)
	fmt.Printf("Gesamt: %.2f €\n", preis+steuer)
	fmt.Printf("Pi (15 Nachkommastellen): %.15f\n", piExakt)

	// --- fmt.Printf — Formatierte Ausgabe ---
	// %d = ganze Zahl (decimal)
	// %f = Dezimalzahl (float)
	// %s = Text (string)
	// %t = Wahrheitswert (bool)
	// %v = "Value" — wähle das passende Format automatisch
	// %T = Typ der Variable

	fmt.Println("\n=== Formatierte Ausgabe ===")
	fmt.Printf("Name: %s, Alter: %d, Größe: %.1f m\n", name, alter, groesse)
	fmt.Printf("Typ von name: %T, Typ von alter: %T\n", name, alter)

	// ==========================================
	// Häufige Fehler – Zusammenfassung (aus dem Video)
	// ==========================================
	//
	// 1. STRINGS: var s string = nil ist VERBOTEN!
	// var s string → s ist leer (""), nicht nil!
	//
	// 2. TYP-KONVERTIERUNG: int + float64 geht NICHT direkt!
	// Immer explizit konvertieren: float64(a) + b
	// Auch verschiedene int-Typen (int + int32) sind inkompatibel.
	// float → int schneidet Nachkommastellen ab (kein Runden!)
	//
	// 3. KONSTANTEN: const ist zur Compile-Zeit fix.
	// Einmal gesetzt → unveränderlich.
	// const x = os.Hostname() geht nicht! Dafür var verwenden.

	fmt.Println("\n Häufige Fehler aus dem Video:")
	fmt.Println(" 1. Strings: var s string = nil → s ist leer (\"\"), nicht nil!")
	fmt.Println(" 2. Typen: int + float64 → immer explizit konvertieren!")
	fmt.Println(" 3. Konstanten: const ist fix zur Compile-Zeit, nicht änderbar ")

	fmt.Println("\n Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// ---------------------------------------------------------------------------
// Übungen
//
// 1. Deklariere Variablen für: deine Lieblingszahl, deine Lieblingsfarbe, dein Geburtsjahr
// 2. Berechne dein Alter in Jahren (ungefähr) und gib es aus
// 3. Definiere eine Konstante für die Lichtgeschwindigkeit (299_792_458 m/s)
// 4. Gib den Satz "Ich heiße [name] und bin [alter] Jahre alt." mit fmt.Printf aus
// 5. Tausche zwei Variablenwerte:
// a := 5; b := 10 → danach: a=10, b=5
// ---------------------------------------------------------------------------
