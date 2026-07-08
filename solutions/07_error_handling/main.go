package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// Übung 1: String in int umwandeln mit Fehlerbehandlung
	werte := []string{"42", "abc", "100", ""}
	for _, v := range werte {
		zahl, err := parseMitFehler(v)
		if err != nil {
			fmt.Printf("parseMitFehler(%q): %v\n", v, err)
		} else {
			fmt.Printf("parseMitFehler(%q) = %d\n", v, zahl)
		}
	}

	// Übung 2: kontoBearbeiten mit errors.Is
	fmt.Println()
	konten := map[string]int{"anna": 100, "unbekannt": 999, "gesperrt": -1}
	for name := range konten {
		err := kontoBearbeiten(name)
		if errors.Is(err, ErrBenutzerNichtGefunden) {
			fmt.Printf("kontoBearbeiten(%q): Benutzer nicht gefunden\n", name)
		} else if errors.Is(err, ErrUngueltigeID) {
			fmt.Printf("kontoBearbeiten(%q): Ungültige ID\n", name)
		} else if err != nil {
			fmt.Printf("kontoBearbeiten(%q): %v\n", name, err)
		} else {
			fmt.Printf("kontoBearbeiten(%q): OK\n", name)
		}
	}

	// Übung 3: defer für Abschlussmeldung
	fmt.Println()
	prozessieren(true)
	prozessieren(false)

	// Übung 4: findPerson mit benutzerdefiniertem Fehler
	fmt.Println()
	_, err := findPerson("Max")
	if err != nil {
		fmt.Printf("findPerson: %v\n", err)
	}
}

// Übung 1: String in int mit eigenen Fehlern
func parseMitFehler(s string) (int, error) {
	if s == "" {
		return 0, errors.New("Eingabe ist leer")
	}
	zahl, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("kann %q nicht in Zahl umwandeln: %w", s, err)
	}
	return zahl, nil
}

// Übung 2: Fehlerarten mit errors.Is unterscheiden
var (
	ErrBenutzerNichtGefunden = errors.New("Benutzer nicht gefunden")
	ErrUngueltigeID          = errors.New("ungültige ID")
)

func kontoBearbeiten(name string) error {
	switch name {
	case "anna":
		return nil
	case "unbekannt":
		return ErrBenutzerNichtGefunden
	default:
		return ErrUngueltigeID
	}
}

// Übung 3: defer für Abschlussmeldung
func prozessieren(fe bool) {
	defer fmt.Println("  -> Funktion prozessieren beendet.")
	if fe {
		fmt.Println("  Fehler aufgetreten!")
		return
	}
	fmt.Println("  Alles OK.")
}

// Übung 4: Benutzerdefinierter Fehler als Struct
type PersonNotFoundError struct {
	Name string
}

func (e *PersonNotFoundError) Error() string {
	return fmt.Sprintf("Person %q nicht gefunden", e.Name)
}

func findPerson(name string) (string, error) {
	return "", &PersonNotFoundError{Name: name}
}
