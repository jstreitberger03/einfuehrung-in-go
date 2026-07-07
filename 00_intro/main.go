// Kapitel 0: Einführung in Go
//
// 🎯 Lernziele:
// - Was ist Go? Warum Go lernen?
// - Go installieren und ein erstes Programm schreiben
// - Den Go Playground kennenlernen
// - Grundlegende Struktur eines Go-Programms verstehen
//
// 📖 Los geht's!

package main

import "fmt"

// main() ist der Einstiegspunkt jedes Go-Programms.
// Wenn du dieses Programm ausführst, startet es hier.
func main() {
	// --- Hallo Welt ---
	// fmt.Println() gibt Text auf der Konsole aus.
	// Der Befehl steht für "Print Line" – also "drucke Zeile".
	fmt.Println("Hallo Welt!")
	fmt.Println("Willkommen zum Go-Kurs! 🎉")

	// --- Ein kleines erstes Programm ---
	// Du kannst direkt Werte ausgeben:
	fmt.Println("2 + 2 =", 2+2)
	fmt.Println("7 * 8 =", 7*8)
	fmt.Println("Hallo" + " " + "Go")

	// --- Kommentare ---
	// Zeilen, die mit // beginnen, sind Kommentare.
	// Sie werden vom Compiler ignoriert und helfen Menschen,
	// den Code zu verstehen.

	// Mehrzeilige Kommentare:
	// /* und */ umschließen einen Block-Kommentar.
	// Alles zwischen /* und */ wird ignoriert.

	// --- Was du gelernt hast ---
	// 1. Ein Go-Programm beginnt immer mit package main und func main()
	// 2. fmt.Println() gibt Text aus
	// 3. Kommentare helfen, Code zu erklären
	// 4. Go kann rechnen und Strings verbinden (+)

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Ändere den Hallo-Welt-Text auf "Hallo, [dein Name]!"
// 2. Gib dein Alter aus: fmt.Println("Ich bin", 25, "Jahre alt")
// 3. Berechne 123 * 456 und gib das Ergebnis aus
// 4. Gib drei Zeilen mit jeweils einem fmt.Println() aus
// ---------------------------------------------------------------------------
//
// 🎬 Weiterführende Ressource (englisch):
// "Learn GO Fast: Full Tutorial" von Alex Mux
// https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=1123s
// Ein kompaktes 1-Stunden-Tutorial, das die Go-Grundlagen wiederholt
// und zeigt, wie man eine API baut. Start bei Minute 18:43.
// ---------------------------------------------------------------------------
