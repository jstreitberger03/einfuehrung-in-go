# Erste Schritte: Theorie vor dem Code

> **Optional.** Lies dieses Kapitel, wenn du zum allerersten Mal programmierst.
> Es beantwortet drei Fragen, die fast jeder Anfänger hat.
> **Kein Code-Schreiben nötig.** Du kannst das Kapitel auch im
> [Go Playground](https://go.dev/play/) ansehen oder einfach auf GitHub lesen.

---

## Willkommen zum "Theorie vor dem Code"-Kapitel!

Dieses Kapitel beantwortet drei Fragen, die fast jeder am Anfang stellt:

1. Was ist überhaupt ein **"Programm"**?
2. Was macht ein **"Compiler"**?
3. Was ist eine **"Variable"**?

Du musst hier NICHTS schreiben. Lies einfach, gerne mit Kaffee. Danach geht es in
[`00_geschichte`](../00_geschichte/) mit der Geschichte der Programmiersprachen weiter.

---

## 1. Was ist ein Programm?

> *Eine genaue Anleitung, wie ein Rezept.*

Stell dir ein Rezept für Kuchen vor:

1. 200 g Mehl in eine Schüssel geben.
2. 100 g Zucker dazu.
3. 2 Eier aufschlagen.
4. Alles 3 Minuten rühren.
5. 30 Minuten backen.

Das ist eine **ANLEITUNG** mit klaren Schritten, in einer Reihenfolge, mit
**ZUTATEN** (Daten) und einem **ERGEBNIS** (Kuchen).

Ein **PROGRAMM** ist genau dasselbe, nur:

| Rezept            | Programm              |
|-------------------|-----------------------|
| Anleitung         | Anweisung an den Computer |
| Zutaten           | Daten / Eingaben       |
| Ergebnis          | Ausgabe auf dem Bildschirm |

**Beispiel aus diesem Kurs:**

|                  |                                          |
|------------------|------------------------------------------|
| **Anweisung**    | Gib "Hallo Welt" aus.                     |
| **Eingabe**      | keine                                    |
| **Ergebnis**     | `"Hallo Welt"` erscheint auf dem Bildschirm |

Mehr nicht. Kein Zauberwerk. Nur eine sehr genaue Anleitung.

---

## 2. Was ist ein Compiler?

> *Ein Übersetzer, der Menschensprache in Maschinensprache umwandelt.*

Computer verstehen nur Nullen und Einsen, die sogenannte **MASCHINENSPRACHE**.

```
"Hallo Welt" als Maschinensprache (Ausschnitt):
    01001000 01100001 01101100 01101100 01101111 …
```

Kein Mensch möchte so schreiben. **Die IDEE:**

- Wir schreiben in einer **LESBAREN** Sprache (Go, Python, …).
- Ein **PROGRAMM** (der Compiler) übersetzt das automatisch.

Stell dir vor: Du schreibst auf Deutsch, der Compiler übersetzt es in die Sprache des Computers.

### Welcher "Übersetzer" nimmt was?

| Befehl              | Was er tut                                                            |
|---------------------|-----------------------------------------------------------------------|
| `go run ./kapitel`  | Nimmt deinen Go-Code und startet das Programm **SOFORT**.              |
| `go build ./kapitel`| Nimmt deinen Code und erzeugt eine `.exe` / Binary-Datei.             |
| `gofmt ./kapitel`   | Formatiert deinen Code hübsch und einheitlich.                        |
| `go vet ./kapitel`  | Prüft auf typische Anfängerfehler.                                    |

### Compiler vs. Interpreter

Du wirst später noch hören: manche Sprachen haben statt eines Compilers einen **"Interpreter"**, der führt Code **Zeile für Zeile** aus, ohne dazwischen eine Binary zu erzeugen.

|           | Compiler (Go)                              | Interpreter (Python)              |
|-----------|--------------------------------------------|-----------------------------------|
| Wie?      | Erst kompilieren, dann starten              | Zeile für Zeile ausführen           |
| Vorteil   | Programm läuft schneller, braucht keine Zusatz-Software beim Start | Kein extra Build-Schritt            |

---

## 3. Was ist eine Variable?

> *Eine beschriftete Schachtel mit Inhalt.*

Stell dir eine beschriftete Kiste vor:

```
┌──────────────────────────┐
│  variable: name          │
│  enthält:    "Anna"       │
└──────────────────────────┘
```

Das ist eine **VARIABLE:**

- Der **NAME** steht außen (`"name"`).
- Der **WERT** steht drin (`"Anna"`).

Was "drin" stehen darf, hängt vom **DATENTYP** ab:

```
┌──────────────────────────┐     ┌──────────────────────┐
│ name       (string)      │     │ alter      (int)     │
│ "Anna"                    │     │ 28                    │
└──────────────────────────┘     └──────────────────────┘
```

### Die vier wichtigsten Datentypen

| Datentyp   | Bedeutung       | Beispiele               |
|------------|-----------------|-------------------------|
| `string`   | Text            | `"Hallo"`, `"Berlin"`, `"x"` |
| `int`      | ganze Zahl      | `1`, `42`, `-7`, `100` |
| `float64`  | Kommazahl       | `3.14`, `1.50`, `-2.5` |
| `bool`     | Ja/Nein         | `true`, `false` |

### Warum Typen wichtig sind

- Du kannst nicht `"Apfel" + 3` rechnen.
- Du kannst nicht `5` mit `"fünf"` mischen.
- Computer brauchen klare Regeln, was womit darf.

### Konstanten

Es gibt auch **KONSTANTEN**, Variablen, deren Wert sich nie ändert.
Beispiel: die Mehrwertsteuer oder die Lichtgeschwindigkeit. Go schreibt dafür
das Wort `const` davor. Mehr dazu in Kapitel [`01_basics`](../01_basics/).

---

## 4. Mini-Glossar zum Mitnehmen

> *Die sieben Wörter, die du heute gehört hast.*

| Begriff            | Kurz erklärt                                                                |
|--------------------|-----------------------------------------------------------------------------|
| **Programm**        | eine genaue Anleitung für einen Computer                                     |
| **Maschinensprache** | die "Sprache" aus Nullen und Einsen, die der Prozessor versteht              |
| **Compiler**        | übersetzt lesbaren Code in Maschinensprache                                  |
| **Interpreter**     | führt Code Schritt für Schritt aus (Python)                                  |
| **Variable**        | eine beschriftete Schachtel mit einem Wert                                   |
| **Datentyp**        | sagt, welche Art von Wert etwas ist (`string`, `int`, …)                      |
| **Konstante**       | eine Variable, deren Wert sich nie ändert                                     |

> Alle 7 Begriffe haben einen Eintrag im großen **Glossar** ([`../GLOSSARY.md`](../GLOSSARY.md)).

---

## 5. War dieses Kapitel hilfreich für dich?

> *Kurze Selbsteinschätzung, wichtig für deinen Lernpfad.*

Wenn du jetzt in einem Satz erklären kannst:

> *"Ein Programm ist eine Anleitung, der Compiler übersetzt sie, und eine Variable ist eine beschriftete Schachtel mit Inhalt."*

……dann bist du **READY** für [`00_geschichte`](../00_geschichte/).

Wenn etwas noch unklar ist:

- Lies den Abschnitt noch einmal in Ruhe.
- Schau dir ein YouTube-Video zum Thema an.
- Frag jemanden, der es dir erklären kann.

Es gibt **KEINEN** Zeitdruck. Lieber einmal verstanden, als dreimal überflogen.

---

## Geschafft! Weiter zu [`00_geschichte`](../00_geschichte/)

> Dort lernst du die Geschichte hinter den Programmiersprachen und **WARUM** ausgerechnet Go für dich interessant sein könnte.

> Wenn du Go schon kennst und direkt loslegen willst, spring einfach zu
> [`00_intro`](../00_intro/) und überspringe die Geschichte. Beides ist okay.

Kapitel abgeschlossen.

---

## Quellcode

Das Go-Programm, das diese Erklärung als Terminal-Output erzeugt, liegt unter
[`00_erste_schritte/main.go`](main.go). Führe es aus mit:

```bash
go run 00_erste_schritte/main.go
```

---

## Übungen (alle optional, alle Theorie, kein Tippen)

1. Erkläre in einem Satz jemand anderem, was ein "Programm" ist. Probiere es wirklich laut aus.
2. Schau dir deinen Computer an: Wo "läuft" das aktuell geöffnete Programm gerade? Auf welchem Prozessor? Was macht der wohl?
3. Nimm eine Tasse, einen Löffel, ein Glas. Welche "Variable" würdest du benennen? Welchen "Datentyp" hat sie?
4. Bonus: Schau im Internet nach "Hello World" in 5 verschiedenen Sprachen. Was fällt dir auf?
