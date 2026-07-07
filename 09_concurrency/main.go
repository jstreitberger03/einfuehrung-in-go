// Kapitel 9: Concurrency — Nebenläufigkeit in Go
//
// 🎯 Lernziele:
// - Goroutinen starten (das Go-Schlüsselwort)
// - Channels zur Kommunikation
// - Buffered vs. unbuffered Channels
// - Select für mehrere Channels
// - WaitGroup zum Warten auf Goroutinen
//
// 📖 Los geht's!

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ==========================================
	// 1. Goroutinen – leichtgewichtige "Threads"
	// ==========================================
	fmt.Println("=== Goroutinen ===")

	// Mit "go" wird eine Funktion parallel ausgeführt
	go sagHallo("Anna")
	go sagHallo("Ben")

	// Kurze Pause, damit die Goroutinen Zeit haben
	// (später ersetzen wir das durch WaitGroup oder Channels)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hauptprogramm fertig")

	// ==========================================
	// 2. Channels – Kommunikation zwischen Goroutinen
	// ==========================================
	fmt.Println("\n=== Channels ===")

	// Ein Channel ist wie ein Rohr: Senden ↔ Empfangen
	nachrichten := make(chan string) // Unbuffered Channel

	// Goroutine sendet eine Nachricht
	go func() {
		nachrichten <- "Hallo vom Channel! 📨"
	}()

	// Hauptprogramm empfängt die Nachricht
	msg := <-nachrichten
	fmt.Println("Empfangen:", msg)

	// ==========================================
	// 3. Buffered Channels
	// ==========================================
	fmt.Println("\n=== Buffered Channels ===")

	// Buffered Channel kann 2 Nachrichten puffern
	zahlen := make(chan int, 2)

	zahlen <- 10
	zahlen <- 20
	// zahlen <- 30 // 💥 WÜRDE BLOCKIEREN (Puffer voll)

	fmt.Println(<-zahlen) // 10
	fmt.Println(<-zahlen) // 20

	// ==========================================
	// 4. Range über Channel
	// ==========================================
	fmt.Println("\n=== Range über Channel ===")

	fibo := make(chan int, 6)

	go func() {
		for i := 0; i < 6; i++ {
			fibo <- i * i
		}
		close(fibo) // Channel schließen – sonst deadlock!
	}()

	for zahl := range fibo {
		fmt.Printf("%d ", zahl)
	}
	fmt.Println()

	// ==========================================
	// 5. Select – auf mehrere Channels warten
	// ==========================================
	fmt.Println("\n=== Select ===")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		chan1 <- "schnell 🏃"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		chan2 <- "langsam 🐢"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("Channel 1:", msg1)
	case msg2 := <-chan2:
		fmt.Println("Channel 2:", msg2)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout ⏰")
	}

	// ==========================================
	// 6. WaitGroup – Goroutinen synchronisieren
	// ==========================================
	fmt.Println("\n=== WaitGroup ===")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Eine Goroutine mehr, auf die gewartet wird
		go arbeiter(i, &wg)
	}

	wg.Wait() // Warte, bis alle fertig sind
	fmt.Println("Alle Arbeiter fertig! ✅")

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// --- Hilfsfunktionen ---

func sagHallo(name string) {
	fmt.Printf("Hallo, %s! (von Goroutine 👋)\n", name)
}

func arbeiter(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal: Diese Goroutine ist fertig
	fmt.Printf("Arbeiter %d startet...\n", id)
	time.Sleep(time.Duration(id*50) * time.Millisecond)
	fmt.Printf("Arbeiter %d fertig ✅\n", id)
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Starte 5 Goroutinen, die gleichzeitig "Hallo von #i" ausgeben
// 2. Schreibe ein Programm mit einem buffered Channel (Puffer = 3)
// 3. Erstelle einen Worker-Pool: 3 Goroutinen lesen von einem Channel
//    und verarbeiten 10 Aufgaben
// 4. Schreibe ein Programm, das mit select auf einen Channel und
//    einen Timeout wartet
// 5. Simuliere ein Rennen: Zwei Goroutinen laufen parallel, wer zuerst
//    5 Punkte erreicht, gewinnt
// ---------------------------------------------------------------------------
