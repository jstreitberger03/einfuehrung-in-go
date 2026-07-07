# Glossar — Einführung in Go

Ein kompaktes Nachschlagewerk für wichtige Go-Begriffe aus dem Kurs.

---

## A

**Array**
: Eine Datenstruktur mit fester Länge. In Go: `[5]int{1, 2, 3, 4, 5}`. Arrays haben eine feste Größe zur Compile-Zeit.

---

## C

**Channel**
: Ein Kommunikationskanal zwischen Goroutinen. Ermöglicht sicheres Senden und Empfangen von Werten (`chan int`).

**Compiler**
: Übersetzt Go-Code in Maschinensprache. `go build` kompiliert das Programm.

**Concurrency (Nebenläufigkeit)**
: Die Fähigkeit, mehrere Aufgaben unabhängig voneinander auszuführen. In Go durch Goroutinen und Channels realisiert.

---

## D

**Defer**
: Ein Schlüsselwort, das die Ausführung einer Funktion bis zum Ende der umgebenden Funktion verschiebt. Ideal für Aufräumarbeiten (z.B. Dateien schließen).

**Dependency**
: Ein externes Paket, von dem dein Modul abhängt. Verwaltet über `go.mod`.

---

## E

**Error**
: Der eingebaute Interface-Typ für Fehlerbehandlung. Jeder Fehler implementiert die `Error()` Methode.

**Exportierter Name**
: Ein Name, der mit Großbuchstaben beginnt und außerhalb des Pakets sichtbar ist.

---

## G

**Go Module**
: Eine Sammlung von Go-Paketen mit einer gemeinsamen Versionsverwaltung. Definiert in `go.mod`.

**GOPATH**
: Der Arbeitsbereich für Go-Code (heute meist durch Module ersetzt).

**Goroutine**
: Ein leichtgewichtiger, paralleler Ausführungsfaden. Gestartet mit dem Schlüsselwort `go`.

---

## I

**Interface**
: Ein Satz von Methoden-Signaturen. Typen implementieren Interfaces implizit — kein explizites `implements` nötig.

---

## M

**Map**
: Eine Sammlung von Schlüssel-Wert-Paaren. Deklaration: `map[string]int{"Apfel": 5}`

**Methode**
: Eine Funktion mit einem Empfänger (`receiver`). Ermöglicht objektorientierte Programmierung in Go.

**Module**
: Siehe Go Module.

---

## N

**nil**
: Der Null-Wert für Pointer, Interfaces, Maps, Slices, Channels und Funktionen. Vergleichbar mit `null` in anderen Sprachen.

---

## P

**Package (Paket)**
: Eine Sammlung von `.go`-Dateien im selben Verzeichnis. Jede Go-Datei gehört zu einem Paket.

**Pointer**
: Speicheradresse einer Variable. Deklariert mit `*T` und erstellt mit `&`.

---

## R

**Range**
: Ein Schlüsselwort zum Iterieren über Slices, Maps, Strings oder Channels. Beispiel: `for i, v := range slice { }`

**Receiver (Empfänger)**
: Der Typ, an den eine Methode gebunden ist. Steht zwischen `func` und dem Methodennamen.

---

## S

**Select**
: Ein Kontrollstruktur, die auf mehrere Channel-Operationen wartet. Ähnlich wie `switch`, aber für Channels.

**Slice**
: Ein dynamisch wachsender Ausschnitt eines Arrays. Flexibler als Arrays. Deklaration: `[]int{1, 2, 3}`

**Struct**
: Ein zusammengesetzter Datentyp, der mehrere Felder bündeln kann. Beispiel: `type Person struct { Name string; Alter int }`

---

## T

**Table-Driven Test**
: Ein Testmuster in Go, bei dem Testfälle in einer Liste (Tabelle) definiert werden.

**Type Inference (Typ-Inferenz)**
: Die Fähigkeit des Compilers, den Typ einer Variable aus dem Wert abzuleiten. Ermöglicht `x := 42` statt `var x int = 42`.

---

## Z

**Zero Value**
: Der automatische Startwert einer Variable ohne explizite Initialisierung:
- `0` für Zahlen
- `false` für bool
- `""` für Strings
- `nil` für Pointer, Slices, Maps, etc.
