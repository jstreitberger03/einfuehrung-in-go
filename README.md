# Einführung in Go

[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8.svg)](https://go.dev/)
[![Go Playground](https://img.shields.io/badge/spielen%20im-Go%20Playground-00ADD8.svg)](https://go.dev/play/)
[![Lizenz](https://img.shields.io/badge/lizenz-MIT-green.svg)](LICENSE)

Ein umfassender Go-Kurs für Anfänger — vom ersten `fmt.Println("Hallo")` bis zu Concurrency, HTTP-Servern und eigenen CLI-Tools.

> [Setup-Anleitung →](SETUP.md) · [Glossar →](GLOSSARY.md) · [FAQ →](FAQ.md) · [Mitmachen →](CONTRIBUTING.md)

---

## 🚀 Schnellstart

| Methode | Dauer | Voraussetzung |
|---------|-------|---------------|
| **[Go Playground](https://go.dev/play/)** | 0 min | Nur ein Browser |
| **[Lokale Installation](SETUP.md)** | 5 min | Go installiert |

Für den Sofortstart ohne Installation: [Go Playground](https://go.dev/play/) öffnen und loslegen!

Für eine geführte Installation siehe die **[Setup-Anleitung](SETUP.md)**.

---

## 📚 Kursinhalt

### Teil 1: Grundlagen
- **00_intro** — **Starte hier**: Was ist Go? Installation, Go Playground, erstes Programm
- **01_basics** — **Kernkonzepte**: Variablen, Datentypen, Konstanten, Ein-/Ausgabe mit `fmt`
- **02_control_flow** — **Entscheidungen treffen**: If/Else, Switch, For-Schleifen
- **03_functions** — **Funktionen**: Definition, Mehrfachrückgaben, benannte Rückgabewerte
- **04_data_structures** — **Daten organisieren**: Arrays, Slices, Maps, Range

### Teil 2: Fortgeschrittene Konzepte
- **05_structs_interfaces** — **Typen selbst definieren**: Structs, Methoden, Interfaces
- **06_pointers** — **Referenzen verstehen**: Pointer, nil, Stack vs. Heap
- **07_error_handling** — **Fehler behandeln**: error-Interface, panic/recover, defer
- **08_packages_modules** — **Code organisieren**: Pakete, go mod, Exportierte Namen

### Teil 3: Praxis & Concurrency
- **09_concurrency** — **Nebenläufigkeit**: Goroutinen, Channels, Select
- **10_file_io** — **Dateien & Daten**: Dateien lesen/schreiben, JSON, CSV
- **11_testing** — **Testen**: Unit-Tests, Table-Driven Tests, Benchmarks
- **12_projects** — **Mini-Projekte**: CLI-Tool, HTTP-Server, Todo-Liste

---

## 📖 Für absolute Anfänger

1. Starte mit `00_intro/` — lies die Kommentare im Code und führe `go run .` aus
2. Experimentiere! Ändere Werte und beobachte die Ausgabe
3. Wenn etwas nicht funktioniert: `go build ./...` zeigt Compiler-Fehler
4. Führe die Tests mit `go test ./...` aus, um dein Wissen zu prüfen
5. Schlage unbekannte Begriffe im **[Glossar](GLOSSARY.md)** nach
6. Versuche die Übungen am Ende jedes Kapitels

> 🔗 **Ohne Go-Installation lernen?**  
> Kopiere Code-Beispiele in den **[Go Playground](https://go.dev/play/)** — alles läuft direkt im Browser!

---

## 🛠️ Editor (VS Code empfohlen)

Installiere diese Erweiterungen für das beste Erlebnis:
- **Go** (`golang.go`) — Sprachunterstützung, Debugging, Linting
- **Go Test Explorer** (`premparihar.gotestexplorer`) — Tests visuell ausführen

Öffne den Kurs-Ordner in VS Code — die Erweiterung wird automatisch vorgeschlagen.

---

## 🎬 Zusätzliche Ressourcen

Ergänzend zu diesem Kurs empfehle ich folgende Video-Ressource:

| Video | Beschreibung |
|-------|-------------|
| [**Learn GO Fast: Full Tutorial**](https://www.youtube.com/watch?v=8uiZC0l4Ajw) von Alex Mux | Ein ca. 1-stündiges, kompaktes Go-Tutorial auf Englisch – ideal als ergänzende Videobegleitung. Deckt die Grundlagen ab und zeigt den Bau einer API. Enthält anschauliche Negativbeispiele (z.B. zu Strings und Konstanten). Die passenden Kapitel im Kurs verweisen darauf mit ⚠️-Hinweisen. |

---

## 📦 Voraussetzungen

- **Go 1.21+** (empfohlen: 1.26)
- Ein Texteditor (VS Code, GoLand, Vim/Neovim)
- Grundlegendes Terminal-Verständnis

---

## 🤝 Mitmachen

Tippfehler gefunden? Idee für eine Verbesserung? Beiträge sind willkommen!

1. Eröffne ein Issue oder eine Diskussion
2. Forke das Repository und erstelle einen Feature-Branch
3. Reiche einen Pull Request ein

Siehe [CONTRIBUTING.md](CONTRIBUTING.md) für Details.

---

## 📄 Lizenz

Dieses Kursmaterial steht unter der MIT-Lizenz zur Verfügung.
