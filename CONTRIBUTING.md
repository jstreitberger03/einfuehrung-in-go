# Mitmachen — Einführung in Go

Danke, dass du zum Kurs beitragen möchtest! 🎉 Hier erfährst du, wie du mitmachen kannst.

---

## Schnelleinstieg

| Was | Wie |
|-----|-----|
| 🐛 **Fehler melden** | [Issue erstellen](https://github.com/jstreitberger03/einfuehrung-in-go/issues/new) |
| 💡 **Verbesserung vorschlagen** | [Issue erstellen](https://github.com/jstreitberger03/einfuehrung-in-go/issues/new) |
| ✍️ **Selbst beitragen** | Fork → Branch → PR (siehe unten) |

---

## Development Setup

### Voraussetzungen

- **Go 1.21+**
- **Git**

### Einrichtung

```bash
# 1. Repository forken und klonen
git clone https://github.com/DEIN_USERNAME/einfuehrung-in-go.git
cd einfuehrung-in-go

# 2. Alles kompiliert?
go build ./...
go test ./...
```

---

## Beitrags-Richtlinien

### 1. Code-Stil
- **`gofmt`** verwenden (wird automatisch von `go fmt ./...` ausgeführt)
- **`go vet ./...`** sollte keine Fehler melden
- Kommentare auf Deutsch (wie der restliche Kurs)
- Klare, verständliche Erklärungen — als würdest du es einem Freund erklären

### 2. Übungsaufgaben
Übungen sollten:
- Klar formuliert sein
- Keine externen Abhängigkeiten haben
- Eine Musterlösung in `solutions/` mitliefern
- Mit `go test` prüfbar sein

### 3. Pull Request Checkliste
- [ ] `go build ./...` läuft fehlerfrei
- [ ] `go test ./...` läuft fehlerfrei
- [ ] `go vet ./...` läuft fehlerfrei
- [ ] `gofmt -s` wurde ausgeführt
- [ ] README bei neuen Kapiteln aktualisiert

---

## Fragen?

- [Issue erstellen](https://github.com/jstreitberger03/einfuehrung-in-go/issues/new)
