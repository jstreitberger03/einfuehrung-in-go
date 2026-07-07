# Einführung in Go

<div align="left">

[![CI](https://github.com/jstreitberger03/einfuehrung-in-go/actions/workflows/ci.yml/badge.svg)](https://github.com/jstreitberger03/einfuehrung-in-go/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jstreitberger03/einfuehrung-in-go)](https://goreportcard.com/report/github.com/jstreitberger03/einfuehrung-in-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/jstreitberger03/einfuehrung-in-go)](go.mod)
[![Lizenz: MIT](https://img.shields.io/badge/Lizenz-MIT-green.svg)](LICENSE)
[![Letzter Commit](https://img.shields.io/github/last-commit/jstreitberger03/einfuehrung-in-go)](https://github.com/jstreitberger03/einfuehrung-in-go/commits/main)
[![Sterne](https://img.shields.io/github/stars/jstreitberger03/einfuehrung-in-go)](https://github.com/jstreitberger03/einfuehrung-in-go/stargazers)

</div>

Ein umfassender Go-Kurs für Anfänger — von „Was ist Programmieren?" bis zu Concurrency, HTTP-Servern und eigenen CLI-Tools.

> [Setup-Anleitung →](SETUP.md) · [Lernpfad →](LEARNING_PATH.md) · [Glossar →](GLOSSARY.md) · [FAQ →](FAQ.md) · [Mitmachen →](CONTRIBUTING.md)

---

## 🎯 Für wen ist dieser Kurs?

Du bist hier richtig, wenn eine dieser Aussagen zutrifft:

> 🟢 **Grün** = du bist Anfänger  
> 🟡 **Gelb** = du hast schon programmiert  
> 🟠 **Orange** = du beherrschst andere Sprachen

- 🟢 **Ich habe noch nie programmiert.** → Starte bei **[`00_erste_schritte/`](00_erste_schritte)** (reine Theorie, ~10 min), gehe dann **[`00_geschichte/`](00_geschichte)** und danach **[`00_intro/`](00_intro)** in Ruhe durch.
- 🟡 **Ich habe schon mal was programmiert (Schule, Hobby).** → Überspringe die Theorie-Kapitel, starte bei **[`00_intro/`](00_intro)**.
- 🟠 **Ich beherrsche eine andere Sprache und will Go lernen.** → Lies die **[`⚠️ Hinweise`](#-wichtige-hinweise-zu-go-vs-anderen-sprachen)** und starte bei **[`01_basics/`](01_basics)**.

---

## 🚀 Schnellstart (30 Sekunden)

Du hast **keine Zeit** für Setup? Los geht's in 30 Sekunden:

1. Öffne **[go.dev/play](https://go.dev/play/)** im Browser
2. Kopiere ein Code-Beispiel aus **[`00_intro/main.go`](00_intro/main.go)** in den Playground
3. Klicke **Run** 🎉

> Kein Download, kein Terminal, keine Installation — Code im Browser ausführen.

Mehr Info (auch für später) in der **[Setup-Anleitung](SETUP.md)**.

---

## 📚 Kursinhalt

### Teil 0: Einstieg — Erst verstehen, dann coden

- **[`00_erste_schritte/`](00_erste_schritte)** — **Optional, wenn du ganz neu bist**: Was ist ein Programm? Was ist ein Compiler? Was ist eine Variable?
- **[`00_geschichte/`](00_geschichte)** — Geschichte der Programmiersprachen, wie Go entstand und warum es die richtige Wahl ist
- **[`00_intro/`](00_intro)** — Was ist Go? Installation, Go Playground, erstes Programm

### Teil 1: Grundlagen

- **[`01_basics/`](01_basics)** — **Kernkonzepte**: Variablen, Datentypen, Konstanten, Ein-/Ausgabe mit `fmt`
- **[`02_control_flow/`](02_control_flow)** — **Entscheidungen treffen**: If/Else, Switch, For-Schleifen
- **[`03_functions/`](03_functions)** — **Funktionen**: Definition, Mehrfachrückgaben, benannte Rückgabewerte
- **[`04_data_structures/`](04_data_structures)** — **Daten organisieren**: Arrays, Slices, Maps, Range

### Teil 2: Fortgeschrittene Konzepte

- **[`05_structs_interfaces/`](05_structs_interfaces)** — **Typen selbst definieren**: Structs, Methoden, Interfaces
- **[`06_pointers/`](06_pointers)** — **Referenzen verstehen**: Pointer, nil, Stack vs. Heap
- **[`07_error_handling/`](07_error_handling)** — **Fehler behandeln**: error-Interface, panic/recover, defer
- **[`08_packages_modules/`](08_packages_modules)** — **Code organisieren**: Pakete, `go mod`, exportierte Namen

### Teil 3: Praxis & Concurrency

- **[`09_concurrency/`](09_concurrency)** — **Nebenläufigkeit**: Goroutinen, Channels, Select
- **[`10_file_io/`](10_file_io)** — **Dateien & Daten**: Dateien lesen/schreiben, JSON, CSV
- **[`11_testing/`](11_testing)** — **Testen**: Unit-Tests, Table-Driven Tests, Benchmarks
- **[`12_projects/`](12_projects)** — **Mini-Projekte**: CLI-Tool, HTTP-Server, Todo-Liste

> 📍 Detaillierter Lernpfad mit Zeitangaben: **[LEARNING_PATH.md](LEARNING_PATH.md)**

---

## ⚠️ Wichtige Hinweise zu Go vs. anderen Sprachen

Go ist **anders** als Java/Python/JavaScript — auch für Umsteiger:

| Konzept | Go | Andere Sprachen |
|---------|-----|-----------------|
| Vererbung | Type Embedding (Composition) | `class Foo extends Bar` |
| Exceptions | `error` als Rückgabewert | `try/catch` |
| Generics | erst seit Go 1.18 (eingeschränkt) | schon länger verfügbar |
| Schleifen | nur `for` (in 5 Varianten) | `for`, `while`, `do…while` |
| Klassen | keine — Structs + Methoden | `class` |
| Build-System | `go build` (eine Binärdatei) | oft komplexer (Gradle, npm, …) |

Details dazu findest du in den jeweiligen Kapiteln (siehe [Zusätzliche Ressourcen](#-zusätzliche-ressourcen)).

---

## 📖 Für absolute Anfänger — Schritt für Schritt

> Wenn du **zum ersten Mal** programmierst, folge diesem Plan.

1. 🆕 **[`00_erste_schritte/`](00_erste_schritte)** lesen — reine Theorie, kein Code-Schreiben nötig
2. 🆕 **[`00_geschichte/`](00_geschichte)** ansehen — verstehen, woher Go kommt
3. 🛠️ In der **[Setup-Anleitung](SETUP.md)** den Go Playground öffnen (kein Download!)
4. 👀 `[00_intro/main.go](00_intro/main.go)` lesen und **kopiere den Code** in den Playground
5. ▶️ Klicke **Run** — du hast gerade dein erstes Programm ausgeführt!
6. ✏️ Ändere einen Text im Code und klicke erneut **Run** — du experimentierst
7. 📘 Weiter zu **[`01_basics/`](01_basics)** und so Stück für Stück vorwärts
8. ❓ Fremde Begriffe? → **[Glossar](GLOSSARY.md)** oder **[FAQ](FAQ.md)**
9. 🏋️ Am Ende jedes Kapitels: Übungen machen (und nicht direkt in die Lösungen schauen!)

> 🔗 **Tipp für zwischendurch:**  
> Du musst nicht alles auf einmal lernen. Pro Tag 30 Minuten reicht völlig.

---

## 🛠️ Editor (VS Code empfohlen)

Wenn du Go später **lokal auf deinem Computer** installierst, hol dir diese Erweiterungen:

- **Go** (`golang.go`) — Sprachunterstützung, Debugging, Linting
- **Go Test Explorer** (`premparihar.gotestexplorer`) — Tests visuell ausführen

Öffne den Kurs-Ordner in VS Code — die Erweiterung wird automatisch vorgeschlagen.

---

## 🎬 Zusätzliche Ressourcen

| Video | Beschreibung |
|-------|-------------|
| [**Learn GO Fast: Full Tutorial**](https://www.youtube.com/watch?v=8uiZC0l4Ajw) von Alex Mux | Ein ca. 1-stündiges, kompaktes Go-Tutorial auf Englisch – ideal als ergänzende Videobegleitung. Deckt die Grundlagen ab und zeigt den Bau einer API. Enthält anschauliche Negativbeispiele (z. B. zu Strings und Konstanten). Die passenden Kapitel im Kurs verweisen darauf mit ⚠️-Hinweisen. |

---

## 📦 Voraussetzungen

Falls du Go **lokal** installieren willst (nicht zwingend nötig):

- **Go 1.24+** (empfohlen: aktuelle stabile Version)
- Ein Texteditor (VS Code, GoLand, Vim/Neovim)
- Grundlegendes Terminal-Verständnis

Anleitung: **[SETUP.md](SETUP.md)**

---

## 🤝 Mitmachen

Tippfehler gefunden? Idee für eine Verbesserung? Beiträge sind willkommen!

1. Eröffne ein Issue oder eine Diskussion
2. Forke das Repository und erstelle einen Feature-Branch
3. Reiche einen Pull Request ein
4. **Warte auf die CI** — der [Status-Badge](https://github.com/jstreitberger03/einfuehrung-in-go/actions/workflows/ci.yml) oben muss grün sein, bevor ein PR gemergt werden kann.

Die CI prüft automatisch:
- ✅ `gofmt -l` — saubere Formatierung
- ✅ `go vet ./...` — keine offiziellen Go-Warnungen
- ✅ `go build ./...` — alle Kapitel kompilieren
- ✅ `go test ./...` — vorhandene Tests laufen
- ✅ `golangci-lint run ./...` — erweiterte Lint-Analyse (mehrere Linter)

Siehe [CONTRIBUTING.md](CONTRIBUTING.md) für Details.

---

## 📄 Lizenz

Dieses Kursmaterial steht unter der MIT-Lizenz zur Verfügung.
