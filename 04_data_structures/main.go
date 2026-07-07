// Kapitel 4: Datenstrukturen — Arrays, Slices, Maps
//
// 🎯 Lernziele:
// - Arrays (feste Größe) vs. Slices (flexibel)
// - Slices: append, Slicing, make
// - Maps: Key-Value-Speicher
// - Range zum Iterieren
//
// 📖 Los geht's!

package main

import "fmt"

func main() {
	// ==========================================
	// 1. Arrays (feste Größe)
	// ==========================================
	fmt.Println("=== Arrays ===")

	// Deklaration mit fester Größe
	var tage [7]string
	tage[0] = "Montag"
	tage[1] = "Dienstag"
	tage[2] = "Mittwoch"
	tage[3] = "Donnerstag"
	tage[4] = "Freitag"
	tage[5] = "Samstag"
	tage[6] = "Sonntag"

	fmt.Println("Alle Tage:", tage)
	fmt.Println("Erster Tag:", tage[0])
	fmt.Println("Letzter Tag:", tage[6])
	fmt.Println("Anzahl Tage:", len(tage))

	// Kurzschreibweise
	fibonacci := [6]int{0, 1, 1, 2, 3, 5}
	fmt.Println("\nFibonacci:", fibonacci)

	// ==========================================
	// 2. Slices (flexibel, dynamisch)
	// ==========================================
	fmt.Println("\n=== Slices ===")

	// Ein Slice wird ohne Größe deklariert
	var noten []int              // nil-Slice (noch leer)
	noten = append(noten, 1)     // Element hinzufügen
	noten = append(noten, 2, 3)  // Mehrere Elemente
	fmt.Println("Noten:", noten)

	// Slice-Literal
	farben := []string{"Rot", "Grün", "Blau"}
	fmt.Println("Farben:", farben)
	fmt.Println("Länge:", len(farben), "Kapazität:", cap(farben))

	// Elemente hinzufügen
	farben = append(farben, "Gelb")
	fmt.Println("Nach append:", farben)

	// --- Slicing: Teil eines Slices ---
	zahlen := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("\nSlicing:")
	fmt.Println("zahlen[1:4]:", zahlen[1:4]) // 20, 30, 40 (Index 1 bis 3)
	fmt.Println("zahlen[:3]:", zahlen[:3])   // 10, 20, 30
	fmt.Println("zahlen[3:]:", zahlen[3:])   // 40, 50, 60

	// --- Slice mit make erstellen ---
	// make([]typ, länge, kapazität)
	slice := make([]int, 3, 5)
	slice[0] = 100
	slice[1] = 200
	slice[2] = 300
	fmt.Println("\nmake-Beispiel:", slice)
	fmt.Println("Länge:", len(slice), "Kapazität:", cap(slice))

	// ==========================================
	// 3. Maps (Key-Value-Speicher)
	// ==========================================
	fmt.Println("\n=== Maps ===")

	// map[key-typ]wert-typ{...}
	alterMap := map[string]int{
		"Anna":  28,
		"Ben":   34,
		"Clara": 22,
	}
	fmt.Println("Alters-Map:", alterMap)

	// Werte abrufen
	fmt.Println("Alter von Anna:", alterMap["Anna"])

	// Werte hinzufügen/ändern
	alterMap["David"] = 31
	fmt.Println("Mit David:", alterMap)

	// Prüfen, ob ein Key existiert
	alter, existiert := alterMap["Eva"]
	if existiert {
		fmt.Println("Alter von Eva:", alter)
	} else {
		fmt.Println("Eva nicht gefunden 😢")
	}

	// Löschen
	delete(alterMap, "Ben")
	fmt.Println("Nach Löschen:", alterMap)

	// --- Nützliche Map-Operationen ---
	fmt.Println("\nAnzahl Einträge:", len(alterMap))

	// ==========================================
	// 4. Range — Über Sammlungen iterieren
	// ==========================================
	fmt.Println("\n=== Range ===")

	// Range über Slice
	fmt.Println("\nFarben (mit Index):")
	for i, farbe := range farben {
		fmt.Printf("%d: %s\n", i, farbe)
	}

	// Range über Map
	fmt.Println("\nPersonen (Name → Alter):")
	for name, a := range alterMap {
		fmt.Printf("%s → %d Jahre\n", name, a)
	}

	// Range mit Ignorieren des Index/Keys
	fmt.Println("\nNur die Werte:")
	for _, farbe := range farben {
		fmt.Println("-", farbe)
	}

	fmt.Println("\n✅ Kapitel abgeschlossen! Probiere die Übungen unten aus.")
}

// ---------------------------------------------------------------------------
// 🏋️ Übungen
//
// 1. Erstelle ein Array mit deinen 3 Lieblingszahlen und gib sie aus
// 2. Erstelle ein leeres Slice und füge 5 Elemente mit append hinzu
// 3. Schreibe Code, der die Summe aller Zahlen in einem Slice berechnet
// 4. Erstelle eine Map "stadtEinwohner" mit 5 Städten und Einwohnerzahlen
// 5. Schreibe eine Funktion, die prüft, ob ein Wert in einem Slice existiert
// 6. Erstelle einen Zähler: Zähle, wie oft jeder Buchstabe in "Hello World" vorkommt
// ---------------------------------------------------------------------------
