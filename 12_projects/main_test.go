package main

import "testing"

func TestTodoHinzufuegen(t *testing.T) {
	liste := &TodoListe{datei: "test_todos.json"}
	liste.hinzufuegen("Testaufgabe")
	if len(liste.Aufgaben) != 1 {
		t.Fatalf("erwartet 1 Aufgabe, got %d", len(liste.Aufgaben))
	}
	if liste.Aufgaben[0].Text != "Testaufgabe" {
		t.Errorf("Text = %q, want %q", liste.Aufgaben[0].Text, "Testaufgabe")
	}
	if liste.Aufgaben[0].ID != 1 {
		t.Errorf("ID = %d, want 1", liste.Aufgaben[0].ID)
	}
	// Cleanup (temp file not written in this test)
}

func TestTodoErledigen(t *testing.T) {
	liste := &TodoListe{datei: "test_todos.json"}
	liste.hinzufuegen("Aufgabe 1")
	liste.hinzufuegen("Aufgabe 2")
	liste.erledigen(1)
	if !liste.Aufgaben[0].Erledigt {
		t.Error("Aufgabe 1 sollte erledigt sein")
	}
	if liste.Aufgaben[1].Erledigt {
		t.Error("Aufgabe 2 sollte nicht erledigt sein")
	}
}

func TestTodoLoeschen(t *testing.T) {
	liste := &TodoListe{datei: "test_todos.json"}
	liste.hinzufuegen("A")
	liste.hinzufuegen("B")
	liste.loeschen(1)
	if len(liste.Aufgaben) != 1 {
		t.Errorf("erwartet 1 Aufgabe nach Löschen, got %d", len(liste.Aufgaben))
	}
	if liste.Aufgaben[0].Text != "B" {
		t.Errorf("verbleibende Aufgabe = %q, want %q", liste.Aufgaben[0].Text, "B")
	}
}

func TestTodoIDAutoIncrement(t *testing.T) {
	liste := &TodoListe{datei: "test_todos.json"}
	liste.hinzufuegen("Erste")
	liste.hinzufuegen("Zweite")
	liste.hinzufuegen("Dritte")
	if liste.Aufgaben[0].ID != 1 || liste.Aufgaben[1].ID != 2 || liste.Aufgaben[2].ID != 3 {
		t.Errorf("IDs sollten 1,2,3 sein, got %d,%d,%d",
			liste.Aufgaben[0].ID, liste.Aufgaben[1].ID, liste.Aufgaben[2].ID)
	}
}
