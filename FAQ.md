# FAQ — Einführung in Go

Häufige Probleme und ihre Lösungen. Steht etwas nicht hier? [Erstelle ein Issue](https://github.com/julianstreitberger/einfuehrung-in-go/issues).

---

## Installation & Setup

### ❓ „go: command not found"

Go ist nicht installiert oder nicht im PATH.

**Lösung:**
```bash
# macOS
brew install go

# Ubuntu/Debian
sudo apt install golang-go

# Oder Installer von https://go.dev/dl/
```

Nach der Installation Terminal neu starten.

---

### ❓ „package main is not a package"

Du hast versucht, `go run .` im falschen Verzeichnis auszuführen.

**Lösung:** Stelle sicher, dass du **im Kapitel-Verzeichnis** bist:
```bash
cd einfuehrung-in-go/00_intro
go run .
```

---

### ❓ „undefined: xyz"

Der Compiler kennt die Variable oder Funktion nicht.

**Häufige Ursachen:**
- Tippfehler im Namen (Go unterscheidet Groß-/Kleinschreibung!)
- Variable wurde nicht deklariert
- Falscher Paketname beim Import

---

## Go-Konzepte

### ❓ Warum `:=` statt `=`?

- `:=` deklariert **und** initialisiert eine neue Variable: `name := "Anna"`
- `=` weist einer **bereits existierenden** Variable einen Wert zu: `name = "Ben"`

---

### ❓ Was bedeutet `_` (underscore) in Go?

`_` ist der **Blank Identifier**. Er ignoriert Werte, die eine Funktion zurückgibt:
```go
sum, _ := divide(10, 3) // ignoriere den zweiten Rückgabewert
```

---

### ❓ „declared but not used" — warum dieser Fehler?

Go erlaubt keine ungenutzten Variablen. Entweder:
- Verwende die Variable, oder
- Ersetze sie mit `_`

```go
// ❌ Fehler: x declared but not used
x := 42

// ✅ Richtig
x := 42
fmt.Println(x)

// ✅ Oder ignoriere bewusst
_ = 42
```

---

### ❓ Wie vergleiche ich Fehler?

```go
// ✅ Richtig (mit ==)
if err == io.EOF {
    // Dateiende erreicht
}

// ✅ Oder mit Fehlerprüfung
if err != nil {
    // Fehler behandeln
}
```

---

### ❓ panic: runtime error: index out of range

Du greifst auf ein Slice-Element zu, das es nicht gibt:
```go
s := []int{1, 2, 3}
fmt.Println(s[5]) // ❌ Index 5 gibt es nicht (0-2)
```

**Lösung:** Prüfe die Länge vor dem Zugriff.

---

## Tests

### ❓ „no test files"

`go test` findet keine Testdateien. Stelle sicher, dass:
- Die Datei auf `_test.go` endet
- Die Funktion mit `Test` beginnt: `func TestXxx(t *testing.T)`

---

### ❓ „go test" zeigt nichts an?

Tests sind erfolgreich durchgelaufen! Mit `-v` siehst du Details:
```bash
go test -v ./...
```

---

## Module & Importe

### ❓ „could not import ..."

Ein importiertes Paket wurde nicht gefunden.

**Lösung:**
```bash
go mod tidy  # Lädt fehlende Abhängigkeiten
```

---

## Hilfe finden

- 🐛 Bug gefunden? [Issue erstellen](https://github.com/julianstreitberger/einfuehrung-in-go/issues)
- 📖 Offizielle [Go-Dokumentation](https://go.dev/doc/)
- 💬 [Go Forum](https://forum.golangbridge.org/)
- 🤖 [Go by Example](https://gobyexample.com/) — Beispiele auf Deutsch/Englisch
