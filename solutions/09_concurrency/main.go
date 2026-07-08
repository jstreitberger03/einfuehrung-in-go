package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Übung 1: 5 Goroutinen gleichzeitig
	fmt.Println("=== Übung 1: 5 Goroutinen ===")
	var wg1 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg1.Add(1)
		go func(id int) {
			defer wg1.Done()
			fmt.Printf("Hallo von #%d\n", id)
		}(i)
	}
	wg1.Wait()

	// Übung 2: Buffered Channel (Puffer = 3)
	fmt.Println("\n=== Übung 2: Buffered Channel ===")
	ch := make(chan string, 3)
	ch <- "Nachricht 1"
	ch <- "Nachricht 2"
	ch <- "Nachricht 3"
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

	// Übung 3: Worker-Pool
	fmt.Println("\n=== Übung 3: Worker-Pool ===")
	aufgaben := make(chan int, 10)
	var wg3 sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg3.Add(1)
		go worker(w, aufgaben, &wg3)
	}
	for i := 1; i <= 10; i++ {
		aufgaben <- i
	}
	close(aufgaben)
	wg3.Wait()

	// Übung 4: Select mit Timeout
	fmt.Println("\n=== Übung 4: Select mit Timeout ===")
	ergebnis := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ergebnis <- "Daten erhalten"
	}()
	select {
	case msg := <-ergebnis:
		fmt.Println("Empfangen:", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout! Keine Daten in 50ms")
	}

	// Übung 5: Rennen zwischen zwei Goroutinen
	fmt.Println("\n=== Übung 5: Rennen ===")
	punkte := make(chan string)
	go laeufer("Läufer A", punkte)
	go laeufer("Läufer B", punkte)
	gewinner := <-punkte
	fmt.Printf("Gewinner: %s!\n", gewinner)
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("  Worker %d verarbeitet Aufgabe %d\n", id, job)
		time.Sleep(10 * time.Millisecond)
	}
}

func laeufer(name string, gewinner chan<- string) {
	punkte := 0
	for punkte < 5 {
		time.Sleep(20 * time.Millisecond)
		punkte++
		fmt.Printf("  %s: %d Punkte\n", name, punkte)
	}
	gewinner <- name
}
