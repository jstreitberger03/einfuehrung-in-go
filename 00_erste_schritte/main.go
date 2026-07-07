// Kapitel 0: Erste Schritte — Theorie vor dem Code
//
// 🎯 Lernziele (theoretisch, kein Schreiben von Code):
// - Verstehen, was ein "Programm" eigentlich ist
// - Verstehen, was ein Compiler tut (und was nicht)
// - Verstehen, was eine "Variable" bedeutet — und was ein "Datentyp"
//
// 🟠 OPTIONAL: Wenn du schon eine andere Sprache kennst, kannst du
//    dieses Kapitel überspringen. Wenn du ALLERERSTE Schritte machst,
//    lies es gern — es beantwortet Fragen, die gleich am Anfang kommen.
//
// 📖 Du musst hier NICHTS schreiben und NICHTS installieren.
//    Du kannst das Kapitel im Browser im [Go Playground](https://go.dev/play/) lesen
//    oder einfach auf GitHub.
//
// 🕘 Reihenfolge für Anfänger:
//    00_erste_schritte (THEORIE, optional) → 00_geschichte (HISTORIE) →
//    00_intro (DEIN ERSTES PROGRAMM) → 01_basics (VARIABLEN & CO.) → …

package main

import "fmt"

// Diese Konstanten machen den Output ruhig und lesbar.
const (
	trennlinie = "────────────────────────────────────────────────────────────"
	kaffee     = "☕"
	kasten     = "📦"
	pfeil      = "➜"
	daumen     = "👍"
	glossary   = "📑"
	check      = "✅"
	hinweis    = "💡"
)

// rubrik druckt eine Sektion.
func rubrik(titel, kurz string) {
	fmt.Println()
	fmt.Println(trennlinie)
	fmt.Println("  " + titel + " — " + kurz)
	fmt.Println(trennlinie)
}

func main() {
	// ============================================================
	// Begrüßung — entwarnend, kein Druck.
	// ============================================================
	fmt.Println()
	fmt.Println("  " + kaffee + " Willkommen zum \"Theorie vor dem Code\"-Kapitel!")
	fmt.Println()
	fmt.Println("   Dieses Kapitel beantwortet drei Fragen, die fast jeder am")
	fmt.Println("   Anfang stellt:")
	fmt.Println()
	fmt.Println("      1. Was ist überhaupt ein \"Programm\"?")
	fmt.Println("      2. Was macht ein \"Compiler\"?")
	fmt.Println("      3. Was ist eine \"Variable\"?")
	fmt.Println()
	fmt.Println("   Du musst hier NICHTS schreiben. Lies einfach — gerne mit")
	fmt.Println("   Kaffee. Danach geht es in 00_geschichte mit der Geschichte")
	fmt.Println("   der Programmiersprachen weiter.")
	fmt.Println()

	// ============================================================
	// Frage 1 — Was ist ein Programm?
	// ============================================================
	rubrik(
		"1. Was ist ein Programm?",
		"Eine genaue Anleitung — wie ein Rezept.",
	)
	fmt.Println()
	fmt.Println("   Stell dir ein Rezept für Kuchen vor:")
	fmt.Println()
	fmt.Println("       1. 200 g Mehl in eine Schüssel geben.")
	fmt.Println("       2. 100 g Zucker dazu.")
	fmt.Println("       3. 2 Eier aufschlagen.")
	fmt.Println("       4. Alles 3 Minuten rühren.")
	fmt.Println("       5. 30 Minuten backen.")
	fmt.Println()
	fmt.Println("   Das ist eine ANLEITUNG mit klaren Schritten, in einer Reihenfolge,")
	fmt.Println("   mit ZUTATEN (Daten) und einem ERGEBNIS (Kuchen).")
	fmt.Println()
	fmt.Println("   Ein PROGRAMM ist genau dasselbe — nur:")
	fmt.Println("   • Die \"Anleitung\" steht für einen Computer.")
	fmt.Println("   • Die \"Zutaten\" heißen Daten oder Eingaben.")
	fmt.Println("   • Das \"Ergebnis\" ist die Ausgabe auf dem Bildschirm.")
	fmt.Println()
	fmt.Println("   Beispiel aus diesem Kurs:")
	fmt.Println()
	fmt.Println(`       Anweisung an den Computer: gib "Hallo Welt" aus.`)
	fmt.Println("       Eingabe:    keine")
	fmt.Println(`       Ergebnis:   "Hallo Welt" erscheint auf dem Bildschirm.`)
	fmt.Println()
	fmt.Println("   Mehr nicht. Kein Zauberwerk. Nur eine sehr genaue Anleitung.")
	fmt.Println()

	// ============================================================
	// Frage 2 — Was ist ein Compiler?
	// ============================================================
	rubrik(
		"2. Was ist ein Compiler?",
		"Ein Übersetzer — von Menschensprache zu Maschinensprache.",
	)
	fmt.Println()
	fmt.Println("   Computer verstehen, wie gesagt, nur Nullen und Einsen —")
	fmt.Println("   die sogenannte MASCHINENSPRACHE.")
	fmt.Println()
	fmt.Println(`       "Hallo Welt" als Maschinensprache (Ausschnitt):`)
	fmt.Println("           01001000 01100001 01101100 01101100 01101111 …")
	fmt.Println()
	fmt.Println("   Kein Mensch möchte so schreiben. Die IDEE:")
	fmt.Println("   • Wir schreiben in einer LESBAREN Sprache (Go, Python, …).")
	fmt.Println("   • Ein PROGRAMM (der Compiler) übersetzt das automatisch.")
	fmt.Println()
	fmt.Println("   Stell dir vor: Du schreibst auf Deutsch, der Compiler")
	fmt.Println("   übersetzt es in die Sprache des Computers.")
	fmt.Println()
	fmt.Println("   Welcher \"Übersetzer\" nimmt was?")
	fmt.Println()
	fmt.Println("       go run ./kapitel   — Nimmt deinen Go-Code und startet das Programm SOFORT.")
	fmt.Println("       go build ./kapitel — Nimmt deinen Code und erzeugt eine .exe / Binary-Datei.")
	fmt.Println("       gofmt ./kapitel    — Formatiert deinen Code hübsch und einheitlich.")
	fmt.Println("       go vet ./kapitel   — Prüft auf typische Anfängerfehler.")
	fmt.Println()
	fmt.Println("   Du wirst später noch hören: manche Sprachen haben statt eines")
	fmt.Println("   Compilers einen \"Interpreter\" — der führt Code Zeile für Zeile")
	fmt.Println("   aus, ohne dazwischen eine Binary zu erzeugen. Python arbeitet")
	fmt.Println("   so. Go macht es anders: erst kompilieren, dann starten.")
	fmt.Println("   Vorteil: das fertige Programm läuft schneller und braucht")
	fmt.Println("   keine Zusatz-Software beim Start.")
	fmt.Println()

	// ============================================================
	// Frage 3 — Was ist eine Variable?
	// ============================================================
	rubrik(
		"3. Was ist eine Variable?",
		"Eine beschriftete Schachtel — mit Inhalt.",
	)
	fmt.Println()
	fmt.Println("   Stell dir eine beschriftete Kiste vor:")
	fmt.Println()
	fmt.Println("       ┌──────────────────────────┐")
	fmt.Println("       │  variable: name          │")
	fmt.Println(`       │  enthält:    "Anna"       │`)
	fmt.Println("       └──────────────────────────┘")
	fmt.Println()
	fmt.Println("   Das ist eine VARIABLE:")
	fmt.Println("   • Der NAME steht außen (\"name\").")
	fmt.Println(`   • Der WERT steht drin ("Anna").`)
	fmt.Println()
	fmt.Println(`   Was "drin" stehen darf, hängt vom DATENTYP ab:`)
	fmt.Println()
	fmt.Println("       ┌──────────────────────────┐     ┌──────────────────────┐")
	fmt.Println("       │ name       (string)      │     │ alter      (int)     │")
	fmt.Println(`       │ "Anna"                    │     │ 28                    │`)
	fmt.Println("       └──────────────────────────┘     └──────────────────────┘")
	fmt.Println()
	fmt.Println("   Beispiel für DICH:")
	fmt.Println()
	fmt.Println(`       string   = Text        ("Hallo", "Berlin", "x")`)
	fmt.Println("       int      = ganze Zahl  (1, 42, -7, 100)")
	fmt.Println("       float64  = Kommazahl   (3.14, 1.50, -2.5)")
	fmt.Println("       bool     = Ja/Nein     (true, false)")
	fmt.Println()
	fmt.Println("   Warum Typen wichtig sind:")
	fmt.Println()
	fmt.Println(`       Du kannst nicht "Apfel" + 3 rechnen.`)
	fmt.Println(`       Du kannst nicht 5 mit "fünf" mischen.`)
	fmt.Println("       Computer brauchen klare Regeln, was womit darf.")
	fmt.Println()
	fmt.Println("   Zum Schluss: es gibt auch KONSTANTEN — Variablen, deren")
	fmt.Println("   Wert sich nie ändert. Beispiel: die Mehrwertsteuer oder")
	fmt.Println("   die Lichtgeschwindigkeit. Go schreibt dafür das Wort")
	fmt.Println("   const davor. Mehr dazu in Kapitel 01_basics.")
	fmt.Println()

	// ============================================================
	// Mini-Glossar — auf einen Blick
	// ============================================================
	rubrik(
		"4. Mini-Glossar zum Mitnehmen",
		"Die sieben Wörter, die du heute gehört hast.",
	)
	fmt.Println()
	fmt.Println("   " + kasten + "  Programm        — eine genaue Anleitung für einen Computer")
	fmt.Println("   " + kasten + "  Maschinensprache — die \"Sprache\" Nullen und Einsen, die der Prozessor versteht")
	fmt.Println("   " + kasten + "  Compiler       — übersetzt lesbaren Code in Maschinensprache")
	fmt.Println("   " + kasten + "  Interpreter    — führt Code Schritt für Schritt aus (Python)")
	fmt.Println("   " + kasten + "  Variable       — eine beschriftete Schachtel mit einem Wert")
	fmt.Println("   " + kasten + "  Datentyp       — sagt, welche Art von Wert etwas ist (string, int, …)")
	fmt.Println("   " + kasten + "  Konstante      — eine Variable, deren Wert sich nie ändert")
	fmt.Println()
	fmt.Println("   " + glossary + "  Alle 7 Begriffe haben einen Eintrag im großen Glossar (../GLOSSARY.md):")
	fmt.Println()
	fmt.Println("        • Programm         → #programm")
	fmt.Println("        • Maschinensprache → #maschinensprache")
	fmt.Println("        • Compiler        → #compiler")
	fmt.Println("        • Interpreter     → #interpreter")
	fmt.Println("        • Variable        → #variable")
	fmt.Println("        • Datentyp        → #datentyp")
	fmt.Println("        • Konstante       → #konstante")
	fmt.Println()

	// ============================================================
	// Optional-Check
	// ============================================================
	rubrik(
		"5. War dieses Kapitel hilfreich für dich?",
		"Kurze Selbsteinschätzung — wichtig für deinen Lernpfad.",
	)
	fmt.Println()
	fmt.Println("   Wenn du jetzt in einem Satz erklären kannst:")
	fmt.Println()
	fmt.Println(`      "Ein Programm ist eine Anleitung, der Compiler übersetzt sie,`)
	fmt.Println(`       und eine Variable ist eine beschriftete Schachtel mit Inhalt."`)
	fmt.Println()
	fmt.Println("   ……dann bist du READY für 00_geschichte.")
	fmt.Println()
	fmt.Println("   Wenn etwas noch unklar ist:")
	fmt.Println()
	fmt.Println("      • Lies den Abschnitt noch einmal in Ruhe.")
	fmt.Println("      • Schau dir ein YouTube-Video zum Thema an.")
	fmt.Println("      • Frag jemanden, der es dir erklären kann.")
	fmt.Println()
	fmt.Println("   Es gibt KEINEN Zeitdruck. Lieber einmal verstanden,")
	fmt.Println("   als dreimal überflogen.")
	fmt.Println()

	// ============================================================
	// Abschluss
	// ============================================================
	fmt.Println()
	fmt.Println(trennlinie)
	fmt.Println("  " + daumen + "  Geschafft! Weiter zu [00_geschichte](../00_geschichte/)")
	fmt.Println(trennlinie)
	fmt.Println()
	fmt.Println("   " + pfeil + "  Dort lernst du die Geschichte hinter den Programmiersprachen")
	fmt.Println("       und WARUM ausgerechnet Go für dich interessant sein könnte.")
	fmt.Println()
	fmt.Println("   " + hinweis + "  Wenn du Go schon kennst und direkt loslegen willst,")
	fmt.Println("       spring einfach zu [00_intro](../00_intro/) und überspringe")
	fmt.Println("       die Geschichte. Beides ist okay.")
	fmt.Println()
	fmt.Println(check + " Kapitel abgeschlossen!")
}

// -------------------------------------------------------------------
// 🏋️ Übungen (alle optional, alle Theorie, kein Tippen)
//
// 1. Erkläre in einem Satz jemand anderem, was ein "Programm" ist.
//    Probiere es wirklich laut aus.
// 2. Schau dir deinen Computer an: Wo "läuft" das aktuell geöffnete
//    Programm gerade? Auf welchem Prozessor? Was macht der wohl?
// 3. Nimm eine Tasse, einen Löffel, ein Glas. Welche "Variable"
//    würdest du benennen? Welchen "Datentyp" hat sie?
// 4. Bonus: Schau im Internet nach "Hello World" in 5 verschiedenen
//    Sprachen. Was fällt dir auf?
// -------------------------------------------------------------------
