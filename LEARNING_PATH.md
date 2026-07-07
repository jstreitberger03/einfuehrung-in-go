# Lernpfad — Einführung in Go

Dieser Lernpfad zeigt dir genau, **in welcher Reihenfolge** du den Kurs durcharbeitest — mit Zeitangaben und Tipps, damit du nicht den Überblick verlierst.

> 🟢 = Pflicht (überspringe nichts) · 🟡 = empfohlen · 🟠 = optional / Bonus

---

## 🎯 Wähle dein Tempo

| Lerntyp | Zeit pro Tag | Kursdauer |
|---------|--------------|-----------|
| 🐇 **Schnell** | 2–3 h | ca. 1 Woche |
| 🚶 **Gemütlich** | 30–60 min | ca. 4 Wochen |
| 🐢 **Sehr langsam** | 15–30 min | ca. 2 Monate |

> Egal welches Tempo — **Konsistenz schlägt Intensität.** Täglich 20 Minuten ist besser als einmal 4 Stunden pro Woche.

---

## 📍 Phase 0: Einstieg (Tag 1)

> 🛠️ Lies zuerst **[SETUP.md](SETUP.md)**. Wenn du keine Installation willst: nur den Go Playground nutzen, kein Download nötig.

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 0a | [`00_erste_schritte/`](00_erste_schritte) | 🕐 10 min | 🟠 **Nur wenn du ganz neu bist** |
| 1  | [`00_geschichte/`](00_geschichte)     | 🕐 15 min | 🟢 **Ja, unbedingt** |

**Lernziele Phase 0:**
- Verstehen, was ein Programm / Compiler / Variable überhaupt ist (Kapitel 0a)
- Verstehen, dass Programmiersprachen Werkzeuge sind, keine Zauberei (Kapitel 1)
- Wissen, warum es hunderte Sprachen gibt und *warum* es Go überhaupt gibt

> 🚦 **Ampel-Test bevor Phase 1 startet:** Kannst du in einem Satz erklären, **warum Go entstand**? Wenn ja → weiter.

---

## 📍 Phase 1: Erste Schritte (Tag 2–3)

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 2 | [`00_intro/`](00_intro) | 🕐 30 min | 🟢 |
| 3 | [`01_basics/`](01_basics) | 🕐 60 min | 🟢 |
| 4 | [`02_control_flow/`](02_control_flow) | 🕐 60 min | 🟢 |

**Lernziele:**
- Dein erstes Go-Programm verstehen und schreiben
- Variablen, Datentypen, Konstanten kennen
- Bedingungen und Schleifen benutzen

> 🏋️ **Tipp:** Mach die Übungen am Ende jedes Kapitels. Wirklich. Sie sind klein, aber bringen viel.

---

## 📍 Phase 2: Strukturen verstehen (Tag 4–7)

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 5 | [`03_functions/`](03_functions) | 🕐 60 min | 🟢 |
| 6 | [`04_data_structures/`](04_data_structures) | 🕐 60 min | 🟢 |
| 7 | [`05_structs_interfaces/`](05_structs_interfaces) | 🕐 90 min | 🟢 |

**Lernziele:**
- Funktionen schreiben und aufrufen
- Arrays, Slices und Maps unterscheiden
- Eigene Typen mit Structs und Interfaces definieren

> ⏸️ **Pause empfohlen** nach diesem Block. Lass die Konzepte sacken.

---

## 📍 Phase 3: Fortgeschrittene Themen (Tag 8–14)

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 8 | [`06_pointers/`](06_pointers) | 🕐 45 min | 🟢 |
| 9 | [`07_error_handling/`](07_error_handling) | 🕐 60 min | 🟢 |
| 10 | [`08_packages_modules/`](08_packages_modules) | 🕐 45 min | 🟢 |

**Lernziele:**
- Pointer verstehen (auch wenn sie „unheimlich" wirken)
- Fehler idiomatisch behandeln
- Code in Pakete strukturieren

---

## 📍 Phase 4: Go's Superkräfte (Tag 15–21)

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 11 | [`09_concurrency/`](09_concurrency) | 🕐 90 min | 🟢 |
| 12 | [`10_file_io/`](10_file_io) | 🕐 60 min | 🟢 |
| 13 | [`11_testing/`](11_testing) | 🕐 45 min | 🟢 |

**Lernziele:**
- Goroutinen und Channels benutzen
- Dateien lesen/schreiben, JSON verarbeiten
- Tests schreiben

---

## 📍 Phase 5: Projekte (Tag 22+)

| # | Kapitel | Dauer | Pflicht? |
|---|---------|-------|----------|
| 14 | [`12_projects/`](12_projects) | 🕐 mehrere Stunden | 🟢 |

**Lernziele:**
- Ein CLI-Tool von Grund auf bauen
- Einen HTTP-Server schreiben
- Eine Todo-Liste mit Speicher erstellen

> 🚀 **Bonus nach dem Kurs:** Baue dein eigenes kleines Projekt. Nur so wird Wissen zu Können.

---

## 🧭 Wenn du nicht weiterkommst

| Problem | Lösung |
|---------|--------|
| Ich verstehe ein Kapitel nicht | Lese es nochmal. Warte einen Tag. Lies es nochmal. Oft „klickt" es Tage später. |
| Ich bekomme einen Compiler-Fehler | Lies die Fehlermeldung **langsam** — Go-Fehler sind oft selbsterklärend. |
| Ich weiß nicht, was ein Begriff bedeutet | Schlage im **[Glossar](GLOSSARY.md)** nach. |
| Ich finde keine Lösung | Frage im **[FAQ](FAQ.md)**, im [Go-Forum](https://forum.golangbridge.org/) oder eröffne ein Issue. |
| Ich habe keine Lust mehr | Mach eine Pause. Lieber 2 Tage Pause als aufgeben. |

---

## 🏆 Meilensteine zum Feiern

Mach dir bewusst, was du schon kannst:

- [ ] Phase 0: Ich kann erklären, was ein Programm / Compiler / Variable ist
- [ ] Phase 0: Ich kann erklären, warum es Programmiersprachen gibt
- [ ] Phase 1: Mein erstes Go-Programm läuft 🎉
- [ ] Phase 2: Ich habe eine eigene Funktion mit Mehrfachrückgabe geschrieben
- [ ] Phase 3: Ich habe einen Pointer-Fehler gefunden und behoben
- [ ] Phase 4: Ich habe 2 Goroutinen gleichzeitig laufen lassen
- [ ] Phase 5: Mein HTTP-Server antwortet im Browser 🎉

> 🎉 Wenn du alle Häkchen gesetzt hast: Herzlichen Glückwunsch — du kannst Go!

---

## 🔗 Weiter nach dem Kurs

| Richtung | Lernquelle |
|----------|-----------|
| Web-APIs mit Go | [`net/http` (Standardbibliothek)](https://pkg.go.dev/net/http) |
| Professionelle Web-Apps | [Gin](https://gin-gonic.com/) oder [Echo](https://echo.labstack.com/) |
| Datenbanken | [GORM](https://gorm.io/) oder [sqlx](https://github.com/jmoiron/sqlx) |
| System-Tools & CLI | [Cobra](https://github.com/spf13/cobra) |
| Tieferes Verständnis | ["The Go Programming Language"](https://www.gopl.io/) (Buch) |
| Offizielle Dokumentation | [go.dev/doc/](https://go.dev/doc/) |

---

**Noch Fragen zum Lernpfad?** → [Issue erstellen](https://github.com/jstreitberger03/einfuehrung-in-go/issues)
