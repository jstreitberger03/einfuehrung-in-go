// Paket textutils – Hilfsfunktionen für Textverarbeitung
package textutils

import "strings"

// Begruesse erstellt einen personalisierten Gruß
func Begruesse(name string) string {
	return "Hallo, " + name + "! "
}

// Grossschreiben wandelt einen String in Großbuchstaben um
func Grossschreiben(text string) string {
	return strings.ToUpper(text)
}

// Verketten verbindet mehrere Strings mit einem Trennzeichen
func Verketten(trennzeichen string, teile ...string) string {
	return strings.Join(teile, trennzeichen)
}
