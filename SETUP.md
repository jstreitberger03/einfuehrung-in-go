# Setup-Anleitung: Einführung in Go

> **Du brauchst nichts zu installieren.**
> 99% dieses Kurses laufen direkt im Browser, auch ohne Go auf deinem Computer.
> Lies Schritt 1. Wenn du dort bleibst, bist du fertig mit dem Setup.

---

## Schritt 1: Sofort loslegen (kein Download, 60 Sekunden)

1. Öffne **[go.dev/play](https://go.dev/play/)** in einem neuen Browser-Tab.
2. Lösche den Code, der dort schon steht.
3. Kopiere diesen **drei Zeilen** Code in das Fenster:

   ```go
   package main
   import "fmt"
   func main() { fmt.Println("Hallo, ich lerne gerade Go!") }
   ```

4. Klicke den blauen **Run**-Knopf.
5. Du hast gerade dein erstes Go-Programm ausgeführt.

Das war's. Du bist startklar für die nächsten Kapitel.

> **Was geht im Playground nicht?**
> Datei-Lesen/Schreiben (Kapitel 10), längere `time.Sleep` (Kapitel 9 ab 5 Sekunden).
> Für diese zwei brauchst du Go lokal. Lies weiter, oder überspringe sie erstmal.

**Weiter zu:** [`00_intro/`](../00_intro/) dort schreibst du deine ersten echten Programme.

---

## Schritt 2 (optional): Go auf deinem Computer installieren

> Nur nötig, wenn du regelmäßig offline üben oder Datei-/Netzwerk-Programme schreiben willst.

### Ein-Klick-Setup pro Betriebssystem

| System | Befehl / Aktion |
|--------|-----------------|
| **macOS** (Homebrew vorhanden) | `brew install go` |
| **Windows** | Installationsdatei von [go.dev/dl/](https://go.dev/dl/) laden, doppelklicken, fertig. |
| **Linux** (Ubuntu / Debian / Mint) | `sudo apt install golang-go` |
| **Linux** (Fedora / RHEL) | `sudo dnf install golang` |
| **Linux** (Arch / Manjaro) | `sudo pacman -S go` |
| **Sonstige / maximale Version** | Offizieller Installer: [go.dev/dl/](https://go.dev/dl/) |

### Installation prüfen: ein einziger Befehl

Öffne ein Terminal und tippe:

```bash
go version
```

Du solltest eine Ausgabe wie `go version go1.26 …` sehen. Falls ja: **Install abgeschlossen.**

Wenn stattdessen `command not found` erscheint, siehe [Hilfe unten](#hilfe-es-funktioniert-nicht).

### Kurs herunterladen: zwei Wege

**Mit Git (empfohlen):**
```bash
git clone https://github.com/jstreitberger03/einfuehrung-in-go.git
cd einfuehrung-in-go
```

**Ohne Git (ZIP):** Auf GitHub oben auf **Code, Download ZIP** klicken, ZIP entpacken, im entpackten Ordner weitermachen.

### Erstes Kapitel starten: ein einziger Befehl

```bash
cd 00_geschichte
go run .
```

Du solltest die Willkommensnachricht und die Geschichte sehen.

---

## Schritt 3 (empfohlen): Code-Editor einrichten

Der einfachste Editor für Go-Anfänger ist **VS Code**.

1. **[VS Code herunterladen](https://code.visualstudio.com/)** und installieren.
2. VS Code öffnen.
3. Erweiterung **Go** (`golang.go`) installieren.
4. Den Kurs-Ordner in VS Code öffnen: **File, Open Folder…**.
5. Die Erweiterung fragt, ob sie fehlende Tools installieren soll: **Alles installieren** klicken.

Fertig. Du bekommst ab jetzt Syntax-Hervorhebung, Auto-Vervollständigung und Fehlerprüfung gratis.

---

## Tests ausführen (optional)

```bash
go test ./...              # alle Tests
go test -v ./...           # ausführlich
go test -cover ./...       # mit Abdeckungs-Bericht
```

---

## Nützliche Befehle

| Befehl | Bedeutung |
|--------|-----------|
| `go run .` | Programm ausführen |
| `go build .` | Programm kompilieren (Binary erstellen) |
| `go test ./...` | Alle Tests ausführen |
| `go fmt ./...` | Code automatisch formatieren |
| `go vet ./...` | Auf häufige Fehler prüfen |
| `go mod tidy` | Abhängigkeiten aufräumen |
| `go get <paket>` | Ein externes Paket herunterladen |

---

## Hilfe: es funktioniert nicht?

| Problem | Schnellste Lösung |
|---------|-------------------|
| `go: command not found` (macOS/Linux) | Terminal schließen und neu öffnen. Bei manueller Installation: `export PATH=$PATH:/usr/local/go/bin` in `~/.bashrc` oder `~/.zshrc` ergänzen. |
| `go: command not found` (Windows) | Terminal **komplett schließen** und neu öffnen (PowerShell oder cmd), damit neue Umgebungsvariablen greifen. |
| `package … not in GOROOT` | Du bist im falschen Ordner. Mit `cd einfuehrung-in-go/00_geschichte` ins richtige Verzeichnis wechseln. |
| `undefined: xxx` | Tippfehler? Groß-/Kleinschreibung prüfen. Go unterscheidet exakt. |
| Editor markiert Code rot, aber `go run .` läuft | VS Code: `Cmd/Ctrl + Shift + P` und *Go: Restart Language Server*. |

---

## Weiter nach dem Setup

1. [`00_geschichte/`](../00_geschichte/): die Geschichte hinter Go (lies sie ggf. nochmal)
2. [`00_intro/`](../00_intro/): schreibe dein erstes Programm
3. [Lernpfad](../LEARNING_PATH.md): Kapitel für Kapitel weiter

**Viel Erfolg!**
