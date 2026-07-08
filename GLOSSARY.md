# Glossar: Einführung in Go

Ein kompaktes Nachschlagewerk für wichtige Go-Begriffe und Konzepte aus dem Kurs.

> Neu: **Geschichte und Hintergründe** findest du ganz unten, passend zum Kapitel `00_geschichte`.

---

## A

**Ada Lovelace**
: Britische Mathematikerin (1815–1852). Sie schrieb 1843 den ersten bekannten Algorithmus, den Plan einer Berechnung, und gilt als erste Programmiererin der Welt.

**Algorithmus**
: Eine eindeutige, schrittweise Anleitung zur Lösung eines Problems. Bevor es Computer gab, mussten Menschen Algorithmen per Hand ausführen.

**Array**
: Eine Datenstruktur mit fester Länge. In Go: `[5]int{1, 2, 3, 4, 5}`. Arrays haben eine feste Größe zur Compile-Zeit.

**Assembler / Maschinensprache**
: Die unterste Sprache, die ein Computer direkt versteht. Für Menschen schwer zu lesen. Wurde in den 1950ern von Hochsprachen wie FORTRAN abgelöst.

---

## B

**BASIC**
: Beginner's All-purpose Symbolic Instruction Code (1964). Eine einfache Sprache für Einsteiger, in den 80ern auf Heimcomputern populär.

---

## C

**C**
: Eine kompakte, schnelle Sprache von Dennis Ritchie (1972, Bell Labs). Fundament fast aller modernen Betriebssysteme. Viele spätere Sprachen (C++, Java, Go) orientieren sich an Cs Syntax.

**Channel**
: Ein Kommunikationskanal zwischen Goroutinen. Ermöglicht sicheres Senden und Empfangen von Werten (`chan int`).

**COBOL**
: Common Business-Oriented Language (1959). Eine Sprache für Wirtschaft und Behörden. Wird heute noch in Banken und Versicherungen eingesetzt.

**Compiler**
: Ein Übersetzer, der lesbaren Code (z. B. Go) in Maschinensprache umwandelt, die der Prozessor direkt versteht. In Go ruft `go build` den Compiler auf; `go run` ruft ihn ebenfalls auf und startet das fertige Programm gleich danach. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

**Concurrency (Nebenläufigkeit)**
: Die Fähigkeit, mehrere Aufgaben unabhängig voneinander auszuführen. In Go durch Goroutinen und Channels realisiert.

---

## D

**Defer**
: Ein Schlüsselwort, das die Ausführung einer Funktion bis zum Ende der umgebenden Funktion verschiebt. Ideal für Aufräumarbeiten (z. B. Dateien schließen).

**Dependency**
: Ein externes Paket, von dem dein Modul abhängt. Verwaltet über `go.mod`.

**Datentyp**
: Sagt, welche **Art von Wert** eine Variable enthält, und damit, was man damit tun darf. Beispiele in Go: `string` (Text), `int` (ganze Zahl), `float64` (Kommazahl), `bool` (Ja/Nein). „Apfel" + 3 rechnen geht nicht, weil die Typen nicht passen. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

---

## E

**Error**
: Der eingebaute Interface-Typ für Fehlerbehandlung. Jeder Fehler implementiert die `Error()` Methode.

**Exportierter Name**
: Ein Name, der mit Großbuchstaben beginnt und außerhalb des Pakets sichtbar ist.

---

## F

**FORTRAN**
: Formula Translation (1957). Eine der **ersten Hochsprachen**, geboren aus der Notwendigkeit, wissenschaftliche Berechnungen für jedermann formulierbar zu machen. Wird bis heute in Forschung und Wettervorhersage genutzt.

---

## G

**Garbage Collection**
: Automatische Speicherbereinigung. Go räumt nicht mehr benutzten Speicher selbst auf, anders als in C/C++.

**Go Module**
: Eine Sammlung von Go-Paketen mit einer gemeinsamen Versionsverwaltung. Definiert in `go.mod`.

**Go Playground**
: Ein Online-Editor auf [go.dev/play](https://go.dev/play/), mit dem du Go-Code direkt im Browser ausführen kannst.

**GOPATH**
: Der Arbeitsbereich für Go-Code (heute meist durch Module ersetzt).

**Goroutine**
: Ein leichtgewichtiger, paralleler Ausführungsfaden. Gestartet mit dem Schlüsselwort `go`.

**Hochsprache**
: Eine Programmiersprache, die **menschenlesbar** ist und nicht direkt vom Computer, sondern von einem Compiler/Interpreter übersetzt wird. Beispiel: Go, Python, Java. Gegensatz: Maschinensprache.

---

## I

**IDE**
: Integrated Development Environment: ein Texteditor + Debugger + Tools in einem Programm (z. B. VS Code, GoLand).

**Interface**
: Ein Satz von Methoden-Signaturen. Typen implementieren Interfaces implizit, kein explizites `implements` nötig.

**Interpreter**
: Führt Code **Zeile für Zeile** direkt aus, ohne vorher eine Binary zu erzeugen. Python arbeitet so. Go geht einen anderen Weg: erst kompilieren, dann starten. Vorteil: schnellerer Start, keine Zusatz-Software nötig. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

---

## J

**Java**
: Objektorientierte Sprache von Sun Microsystems (1995). Ziel: einmal schreiben, überall ausführen. Liegt in der Beliebtheit lange weit vorne.

**JavaScript**
: Skriptsprache für Webbrowser (1995). Heute auch auf Servern (Node.js) und fast überall im Web.

---

## K

**Konstante**
: Eine Variable, deren Wert sich **nie ändert**. Beispiele: die Mehrwertsteuer, die Lichtgeschwindigkeit, Pi. In Go steht das Schlüsselwort `const` davor: `const mwst = 0.19`. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

---

## L

**LISP**
: Eine der ältesten Programmiersprachen (1958), stark in der KI-Forschung verwendet.

---

## M

**Map**
: Eine Sammlung von Schlüssel-Wert-Paaren. Deklaration: `map[string]int{"Apfel": 5}`

**Maschinensprache**
: Die Sprache, die ein Computerprozessor direkt versteht: Folgen von Nullen und Einsen. Menschen schreiben kaum noch direkt darin. Hochsprachen wie Go werden erst vom Compiler in Maschinensprache übersetzt, bevor der Prozessor sie ausführen kann. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

**Methode**
: Eine Funktion mit einem Empfänger (`receiver`). Ermöglicht objektorientierte Programmierung in Go.

**Modul**
: Siehe Go Module.

---

## N

**nil**
: Der Null-Wert für Pointer, Interfaces, Maps, Slices, Channels und Funktionen. Vergleichbar mit `null` in anderen Sprachen.

---

## P

**Package (Paket)**
: Eine Sammlung von `.go`-Dateien im selben Verzeichnis. Jede Go-Datei gehört zu einem Paket.

**Paradigma (Programmierparadigma)**
: Grundlegende Programmierweise. Beispiele: imperativ (C), objektorientiert (Java), funktional (LISP/Haskell), nebenläufig (Go).

**Pascal**
: Eine strukturierte Sprache (1970), in den 80ern für die Lehre populär.

**PHP**
: Serverseitige Skriptsprache für Webseiten (1995). Treibt WordPress, Wikipedia und einen großen Teil des Webs.

**Pointer**
: Speicheradresse einer Variable. Deklariert mit `*T` und erstellt mit `&`.

**Programm**
: Eine genaue, geordnete Anleitung für einen Computer, vergleichbar mit einem Rezept: Zutaten (Daten) + Schritte + Ergebnis (Ausgabe auf dem Bildschirm). „Hallo Welt" ist das einfachste mögliche Programm. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

**Python**
: Lesbare Sprache für schnelle Ergebnisse (1991). Beliebt in Datenanalyse, KI und Automatisierung.

---

## R

**Range**
: Ein Schlüsselwort zum Iterieren über Slices, Maps, Strings oder Channels. Beispiel: `for i, v := range slice { }`

**Receiver (Empfänger)**
: Der Typ, an den eine Methode gebunden ist. Steht zwischen `func` und dem Methodennamen.

**Rust**
: Moderne Systemsprache (2010), sehr sicher, aber mit steiler Lernkurve.

---

## S

**Select**
: Eine Kontrollstruktur, die auf mehrere Channel-Operationen wartet. Ähnlich wie `switch`, aber für Channels.

**Slice**
: Ein dynamisch wachsender Ausschnitt eines Arrays. Flexibler als Arrays. Deklaration: `[]int{1, 2, 3}`

**Spaghetti-Code**
: Schlecht strukturierter Code, bei dem der Programmfluss schwer nachvollziehbar ist. Strukturiertes Programmieren (Struktogramme, Funktionen) half ab den 70ern, das zu lösen.

**Struct**
: Ein zusammengesetzter Datentyp, der mehrere Felder bündeln kann. Beispiel: `type Person struct { Name string; Alter int }`

---

## T

**Table-Driven Test**
: Ein Testmuster in Go, bei dem Testfälle in einer Liste (Tabelle) definiert werden.

**Turingmaschine**
: Ein gedankliches Computermodell von Alan Turing (1936). Theoretischer Vorläufer aller heutigen Computer.

**Type Inference (Typ-Inferenz)**
: Die Fähigkeit des Compilers, den Typ einer Variable aus dem Wert abzuleiten. Ermöglicht `x := 42` statt `var x int = 42`.

---

## V

**Variable**
: Eine beschriftete Schachtel mit einem Wert. Der **Name** steht außen, der **Wert** steht drin. Welcher Wert erlaubt ist, bestimmt der Datentyp. In Go: `name := "Anna"` deklariert die Variable `name` und weist ihr den String `"Anna"` zu. Siehe Kapitel [00_erste_schritte](00_erste_schritte).

---

## Z

**Zero Value**
: Der automatische Startwert einer Variable ohne explizite Initialisierung:
- `0` für Zahlen
- `false` für bool
- `""` für Strings
- `nil` für Pointer, Slices, Maps, etc.

---

## Zeitleiste (Bonus)

| Jahr | Ereignis |
|------|----------|
| 1843 | Ada Lovelace beschreibt ersten Algorithmus |
| 1936 | Turingmaschine (Alan Turing) |
| 1957 | FORTRAN |
| 1958 | LISP |
| 1959 | COBOL |
| 1964 | BASIC |
| 1970 | Pascal |
| 1972 | C |
| 1985 | C++ |
| 1991 | Python |
| 1995 | Java, JavaScript, PHP |
| 2009 | **Go (Google)** |
| 2010 | Rust |
| 2014 | Swift |
