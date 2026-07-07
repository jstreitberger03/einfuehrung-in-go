# Setup-Anleitung — Einführung in Go

Willkommen! Diese Anleitung hilft dir, Go für diesen Kurs einzurichten — egal ob Windows, macOS oder Linux.

---

## Schnellstart (2 Optionen)

### Option 1: Go Playground (Keine Installation) 🚀

Der [Go Playground](https://go.dev/play/) ist eine Web-Umgebung, in der du Go-Code direkt im Browser ausführen kannst:

- Keine Installation nötig
- Beispiele aus dem Kurs einfach kopieren und einfügen
- Ideal für den Einstieg

**Einschränkung:** Keine Datei-Ein-/Ausgabe, keine Netzwerkverbindungen, keine unendlichen Schleifen.

### Option 2: Lokale Installation ⚡

#### Schritt 1: Go installieren

| Betriebssystem | Anleitung |
|----------------|-----------|
| **macOS** | `brew install go` |
| **Windows** | Installer von [go.dev](https://go.dev/dl/) herunterladen |
| **Linux** | `sudo apt install golang-go` (Ubuntu/Debian) oder von [go.dev](https://go.dev/dl/) |

Installation prüfen:
```bash
go version
# Sollte anzeigen: go 1.21 oder höher
```

#### Schritt 2: Kurs herunterladen

```bash
git clone https://github.com/julianstreitberger/einfuehrung-in-go.git
cd einfuehrung-in-go
```

#### Schritt 3: Erste Schritte

```bash
# Gehe ins erste Kapitel
cd 00_intro

# Führe den Code aus
go run .
```

---

## 🛠️ Editor-Einrichtung

### VS Code

1. VS Code öffnen
2. Erweiterung **Go** (`golang.go`) installieren
3. Den Kurs-Ordner öffnen: `File → Open Folder...`
4. Die Erweiterung schlägt automatisch fehlende Tools vor — alles installieren

### Weitere Editoren

- **GoLand** — JetBrains IDE mit starkem Go-Support
- **Vim/Neovim** — Mit `vim-go` Plugin
- **Emacs** — Mit `go-mode`

---

## 🧪 Tests ausführen

```bash
# Alle Tests im gesamten Kurs ausführen
go test ./...

# Tests eines bestimmten Kapitels
go test ./01_basics/...

# Mit Coverage-Bericht
go test -cover ./...
```

---

## Häufige Probleme

| Problem | Lösung |
|---------|--------|
| `go: command not found` | Go ist nicht installiert → Siehe Schritt 1 |
| `package xxx is not in GOROOT` | Du bist nicht im richtigen Verzeichnis → `cd` ins Kapitel |
| `undefined: xxx` | Tippfehler? Groß-/Kleinschreibung beachten |
| Modul-Fehler | `go mod tidy` ausführen |

---

## Hilfe benötigt?

- 📖 [Go-Dokumentation](https://go.dev/doc/)
- 🎮 [Go Playground](https://go.dev/play/)
- 💬 [Go Forum](https://forum.golangbridge.org/)
