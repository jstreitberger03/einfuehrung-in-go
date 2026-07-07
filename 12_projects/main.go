// Kapitel 12: Mini-Projekte — Das Gelernte anwenden
//
// 🎯 Lernziele:
// - Ein CLI-Tool bauen (Taschenrechner)
// - Einen einfachen HTTP-Server schreiben
// - Eine Todo-Liste mit Speicher
//
// 📖 Los geht's!

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Mini-Projekte ===")
	fmt.Println()
	fmt.Println("Dieses Kapitel enthält 3 Projekte. Wähle eines aus:")
	fmt.Println()
	fmt.Println("1️⃣  CLI-Taschenrechner")
	fmt.Println("   go run . taschenrechner 5 + 3")
	fmt.Println()
	fmt.Println("2️⃣  HTTP-Server")
	fmt.Println("   go run . server")
	fmt.Println()
	fmt.Println("3️⃣  Todo-Liste")
	fmt.Println("   go run . todo")
	fmt.Println()

	// Projekt-Auswahl
	if len(os.Args) < 2 {
		fmt.Println("Bitte ein Projekt wählen: taschenrechner, server, oder todo")
		return
	}

	switch os.Args[1] {
	case "taschenrechner":
		taschenrechnerCLI()
	case "server":
		httpServer()
	case "todo":
		todoApp()
	default:
		fmt.Printf("Unbekanntes Projekt: %s\n", os.Args[1])
	}
}

// ==========================================
// Projekt 1: CLI-Taschenrechner
// ==========================================

func taschenrechnerCLI() {
	if len(os.Args) != 4 {
		fmt.Println("Verwendung: go run . taschenrechner <zahl1> <op> <zahl2>")
		fmt.Println("Operationen: + - * /")
		return
	}

	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Fehler: erste Zahl ungültig")
		return
	}

	b, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Fehler: zweite Zahl ungültig")
		return
	}

	op := os.Args[3]
	var ergebnis float64

	switch op {
	case "+":
		ergebnis = a + b
	case "-":
		ergebnis = a - b
	case "*":
		ergebnis = a * b
	case "/":
		if b == 0 {
			fmt.Println("Fehler: Division durch Null!")
			return
		}
		ergebnis = a / b
	default:
		fmt.Println("Fehler: unbekannte Operation. Erlaubt: + - * /")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, ergebnis)
}

// ==========================================
// Projekt 2: Einfacher HTTP-Server
// ==========================================

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Gast"
		}
		fmt.Fprintf(w, "Hallo, %s! 👋\n", name)
		fmt.Fprintf(w, "Pfad: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Methode: %s\n", r.Method)
	})

	// JSON-API: /api/begruessung
	http.HandleFunc("/api/begruessung", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Gast"
		}
		antwort := map[string]string{
			"nachricht": fmt.Sprintf("Hallo, %s!", name),
		}
		json.NewEncoder(w).Encode(antwort)
	})

	port := ":8080"
	fmt.Printf("Server läuft auf http://localhost%s\n", port)
	fmt.Println("Öffne http://localhost:8080/ in deinem Browser")
	fmt.Println("Probiere: http://localhost:8080/?name=Anna")
	fmt.Println("API: http://localhost:8080/api/begruessung?name=Ben")
	log.Fatal(http.ListenAndServe(port, nil))
}

// ==========================================
// Projekt 3: Todo-Liste
// ==========================================

type Todo struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Erledigt bool   `json:"erledigt"`
}

type TodoListe struct {
	Aufgaben []Todo `json:"aufgaben"`
	datei    string
}

func todoApp() {
	liste := TodoListe{datei: "todos.json"}
	liste.laden()

	if len(os.Args) < 3 {
		fmt.Println("Verwendung:")
		fmt.Println("  go run . todo list          — Alle Aufgaben anzeigen")
		fmt.Println("  go run . todo add <text>    — Aufgabe hinzufügen")
		fmt.Println("  go run . todo done <id>     — Aufgabe erledigen")
		fmt.Println("  go run . todo delete <id>   — Aufgabe löschen")
		return
	}

	switch os.Args[2] {
	case "list":
		liste.anzeigen()
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Bitte Aufgabentext angeben")
			return
		}
		text := strings.Join(os.Args[3:], " ")
		liste.hinzufuegen(text)
	case "done":
		if len(os.Args) < 4 {
			fmt.Println("Bitte Aufgaben-ID angeben")
			return
		}
		id, _ := strconv.Atoi(os.Args[3])
		liste.erledigen(id)
	case "delete":
		if len(os.Args) < 4 {
			fmt.Println("Bitte Aufgaben-ID angeben")
			return
		}
		id, _ := strconv.Atoi(os.Args[3])
		liste.loeschen(id)
	default:
		fmt.Printf("Unbekannter Befehl: %s\n", os.Args[2])
	}
}

func (l *TodoListe) laden() {
	daten, err := os.ReadFile(l.datei)
	if err != nil {
		return // Keine Datei = leere Liste
	}
	json.Unmarshal(daten, l)
}

func (l *TodoListe) speichern() {
	daten, _ := json.MarshalIndent(l, "", "  ")
	os.WriteFile(l.datei, daten, 0644)
}

func (l *TodoListe) anzeigen() {
	if len(l.Aufgaben) == 0 {
		fmt.Println("📝 Keine Aufgaben vorhanden")
		return
	}
	fmt.Println("📋 Todo-Liste:")
	for _, aufgabe := range l.Aufgaben {
		status := "⬜"
		if aufgabe.Erledigt {
			status = "✅"
		}
		fmt.Printf("  %s [%d] %s\n", status, aufgabe.ID, aufgabe.Text)
	}
}

func (l *TodoListe) hinzufuegen(text string) {
	maxID := 0
	for _, a := range l.Aufgaben {
		if a.ID > maxID {
			maxID = a.ID
		}
	}
	l.Aufgaben = append(l.Aufgaben, Todo{
		ID:   maxID + 1,
		Text: text,
	})
	l.speichern()
	fmt.Printf("✅ Aufgabe hinzugefügt: %s\n", text)
}

func (l *TodoListe) erledigen(id int) {
	for i, a := range l.Aufgaben {
		if a.ID == id {
			l.Aufgaben[i].Erledigt = true
			l.speichern()
			fmt.Printf("✅ Aufgabe erledigt: %s\n", a.Text)
			return
		}
	}
	fmt.Printf("❌ Aufgabe %d nicht gefunden\n", id)
}

func (l *TodoListe) loeschen(id int) {
	for i, a := range l.Aufgaben {
		if a.ID == id {
			l.Aufgaben = append(l.Aufgaben[:i], l.Aufgaben[i+1:]...)
			l.speichern()
			fmt.Printf("🗑️ Aufgabe gelöscht: %s\n", a.Text)
			return
		}
	}
	fmt.Printf("❌ Aufgabe %d nicht gefunden\n", id)
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Erweitere den Taschenrechner um Potenzieren und Wurzel
// 2. Füge dem HTTP-Server eine /api/status-Endpunkt hinzu
// 3. Erweitere die Todo-Liste um Prioritäten (hoch/mittel/niedrig)
// 4. Baue einen Mini-CLI zu "Währungsrechner" (EUR → USD)
// 5. Schreibe einen einfachen Zahlenratespiel-CLI
// ---------------------------------------------------------------------------
