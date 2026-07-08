# Geschichte der Programmiersprachen

> *Wie entstanden Programmiersprachen? Warum gerade Go?*
> Eine kleine Reise durch 12 Sprachen, von 1843 bis heute.
> **Kein Code-Schreiben nötig.** Lesedauer: ca. 5 Minuten.

---

## Schön, dass du da bist!

Dies ist kein Programmier-Test, sondern eine kleine **Geschichte**, erzählt von
Computern, Ideen und Menschen, die vor dir dieselben Fragen gestellt haben.

---

## Wie kommt eine Programmiersprache in die Welt?

> *Jede Sprache entsteht, weil jemand ein Problem hatte.*

Stell dir vor: Computer verstehen nur Nullen und Einsen. Maschinensprache
sieht so aus: `01001000 01100101 01101100 …`

Schon nach kurzer Zeit fragten sich findige Köpfe:

> *"Kann man Computern nicht einfach sagen, was sie tun sollen, in Worten,
> die wir Menschen verstehen?"*

### So sieht "Hallo Welt" in Go aus

```
┌────────────────────────────────────────────────────────────┐
│ So sieht "Hallo Welt" in Go aus ─────────────────────────│
│────────────────────────────────────────────────────────────│
│  CODE  (das, was der Entwickler tippt):                 │
│   fmt.Println("Hallo Welt!")                              │
│────────────────────────────────────────────────────────────│
│  OUTPUT  (das, was der Computer zeigt):                 │
│   Hallo Welt!                                             │
└────────────────────────────────────────────────────────────┘
```

---

## Fünf Epochen, in denen sich die Welt verändert hat

> *Von 1843 bis heute, eine Reise in fünf Etappen.*

| Epoche | Zeitraum | Was geschah |
|--------|----------|-------------|
| Die Anfänge | 1843–1940 | Ideen ohne Computer, Berechnungspläne auf Papier |
| Erste Hochsprachen | 1950–1969 | Maschinen werden leistungsfähiger, Wörter ersetzen die Zahlenwüste |
| Strukturierte Programmierung | 1970–1989 | "Spaghetti-Code" war ein echtes Problem. Geordnete Strukturen und das legendäre C |
| Web-Ära | 1990–2009 | Das Internet verändert alles, schnell schreiben ist wichtiger als schnell laufen |
| Cloud und Moderne | seit 2009 | Millionen Aufgaben gleichzeitig, neue Sprachen (Go, Rust, Kotlin) antworten darauf |

---

## Die wichtigsten Sprachen: wofür sie gut sind

> *Du musst keine einzige davon heute lernen.*

Diese Übersicht soll dir nur ein Gefühl dafür geben, warum es so viele
Sprachen gibt, und warum jede ihre Daseinsberechtigung hat.

| Sprache | Jahr | Kategorie | Hello World | Stärke | Schwäche | Heute |
|---------|------|-----------|-------------|--------|----------|-------|
| **FORTRAN** | 1957 | Wissenschaft | `PRINT *, 'Hello'` | superschnell für numerische Berechnungen | veraltete Syntax für moderne Augen | Wetter- und Klimamodelle, Physik-Simulationen |
| **LISP** | 1958 | Forschung | `(print "Hello")` | elegante Mathematik, Code = Daten | viele Klammern schrecken ab | KI-Forschung, Sprachverarbeitung |
| **COBOL** | 1959 | Wirtschaft | `DISPLAY 'Hello'.` | extrem gut für Buchhaltung und Reports | sehr geschwätzig, schwer zu modernisieren | Banken, Versicherungen, Behörden |
| **BASIC** | 1964 | Bildung | `10 PRINT "Hello"` | sehr einfach zu lernen | untauglich für große Projekte | Heimcomputer in den 80ern, heute Nische |
| **Pascal** | 1970 | Lehre | `writeln('Hello');` | klar und didaktisch | kaum noch in der Industrie | Schulen, historisches Delphi |
| **C** | 1972 | Systeme | `printf("Hello\n");` | extrem schnell, volle Speicherkontrolle | Fehler im Speicher sind berüchtigt | Betriebssysteme, Datenbanken, Go selbst |
| **C++** | 1985 | Systeme | `cout << "Hello" << endl;` | C + Objektorientierung + Templates | viele Wege, es zu nutzen, viele Wege, Fehler zu machen | Spiele, Browser, Bürosoftware |
| **Python** | 1991 | Allzweck | `print('Hello')` | sehr lesbar, riesige Bibliotheken | langsam, schwer für Mobile/Spiele | Data Science, KI, Skripte |
| **Java** | 1995 | Enterprise | `System.out.println("Hello");` | einmal schreiben, überall laufen | viel Speicher und Boilerplate | Große Firmen, Android-Apps |
| **JavaScript** | 1995 | Web | `console.log('Hello');` | läuft in jedem Browser | chaotisch gewachsen, viele Frameworks | Web-Frontends, mit Node.js auch Server |
| **PHP** | 1995 | Server-Web | `echo 'Hello';` | schnell zur lauffähigen Webseite | unübersichtlicher Code bei großen Projekten | WordPress, Wikipedia, große Teile des Webs |
| **Go** | 2009 | Cloud und Server | `fmt.Println("Hello")` | schnell wie C, einfach wie Python, sicher wie Java | weniger Bibliotheken als Python/JS, keine Klassen-Vererbung | Cloud, Microservices, CLI, DevOps |

---

## Es gibt keine "beste Sprache", nur das beste Werkzeug

> *Sprachen sind wie Werkzeuge in einer Werkstatt.*

Ein Hammer ist nicht "besser" als eine Säge. Es kommt darauf an, was du bauen willst.

Wer eine App für 10.000 Menschen gleichzeitig schreibt, trifft andere
Entscheidungen als jemand, der eine Webseite für einen kleinen Verein baut.
Deshalb gibt es heute über 200 Sprachen, und jede Woche neue.
Das ist kein Chaos, sondern Vielfalt.

---

## Und warum schauen wir uns jetzt Go an?

> *Go wurde 2009 bei Google geboren, aus einem echten Problem.*

Google hatte Anfang der 2000er ein gewaltiges Problem:

- **Tausende** Entwickler schreiben Code.
- **Millionen** Nutzer sprechen gleichzeitig mit den Servern.
- Der Code-Check dauerte bei großen Projekten **30+ Minuten**.
- Die Computer verbrauchten viel zu viel Strom und Speicher.

Drei erfahrene Entwickler machten sich daran, eine Sprache zu entwerfen, die
**ALLE** diese Probleme adressiert:

- **Rob Pike** (Mit-Erfinder von Unix)
- **Ken Thompson** (Mit-Erfinder von Unix und C)
- **Robert Griesemer** (Schöpfer der Java-HotSpot-VM)

Ihr Ziel war nicht "die perfekte Sprache", sondern: *die richtige Sprache für
die Aufgaben von heute.*

### So kurz ist ein komplettes Go-Programm

```
┌────────────────────────────────────────────────────────────┐
│ So kurz ist ein komplettes Go-Programm ──────────────────│
│────────────────────────────────────────────────────────────│
│  CODE  (das, was der Entwickler tippt):                 │
│   package main                                            │
│   import "fmt"                                            │
│   func main() { fmt.Println("Hallo, Geschichte!") }       │
│────────────────────────────────────────────────────────────│
│  OUTPUT  (das, was der Computer zeigt):                 │
│   Hallo, Geschichte!                                      │
└────────────────────────────────────────────────────────────┘
```

---

## Was Go anders macht als andere Sprachen

> *Sechs Eigenschaften, die du beim Lernen wiedererkennen wirst.*

1. **Einfach zu lesen, einfach zu lernen**
   Go hat nur 25 Schlüsselwörter. C hat über 30. Wer Go liest, versteht es, auch
   wenn er es nicht selbst schreibt.

2. **Superschnelle Kompilierung**
   Große Projekte wie Kubernetes (~3 Mio. Zeilen) bauen in unter 60 Sek.
   Vergleich: C++-Projekte ähnlicher Größe brauchen Minuten.

3. **Eingebaute Nebenläufigkeit (Concurrency)**
   Tausende "Goroutinen" starten in einer einzigen Datei. Wer 10.000 Aufgaben
   gleichzeitig erledigen muss: Go ist oft die einfachste Sprache dafür.

4. **Eine Datei, ein Programm**
   Go baut eine einzelne `.exe`-Datei (oder Linux-Binary). Kein Runtime,
   keine VM, keine Frameworks-Liste. Starten, fertig.

5. **Speicher wird automatisch aufgeräumt (Garbage Collection)**
   In C/C++ muss man Speicher selbst freigeben und vergisst es leicht.
   In Java kostet das viel Performance. Go macht beides besser.

6. **Eine Sprache, ein Stil**
   Es gibt nur einen verbreiteten Weg, Code zu formatieren. `go fmt`
   formatiert deinen Code automatisch korrekt. Streit über Einrückung?
   Gibt es in Go nicht.

---

## Wo du Go heute überall im Einsatz siehst

> *Auch wenn du es nicht merkst.*

### Berühmte Open-Source-Projekte in Go

| Projekt | Beschreibung |
|---------|-------------|
| **Docker** | Container, die deine Apps starten |
| **Kubernetes** | Verwaltet tausende Container |
| **Terraform** | Beschreibt Cloud-Infrastruktur als Code |
| **Hugo** | Statische Webseiten |
| **Caddy** | Webserver mit automatischem HTTPS |
| **Prometheus** | Monitoring für moderne Systeme |
| **CockroachDB** | Verteilte SQL-Datenbank |

### Große Firmen, die Go produktiv einsetzen

| Firma | Einsatz |
|-------|---------|
| **Google** | Ursprung und viele interne Tools |
| **Uber** | Geokoordinaten und Matching in Echtzeit |
| **Twitch** | Video-Pipeline für Live-Streams |
| **Stripe** | Teile der Zahlungs-Infrastruktur |
| **Cloudflare** | DNS, Edge-Computing, Sicherheit |
| **Dropbox** | Magischer Datei-Sync |

Go ist keine "Geheimtipp"-Sprache mehr, sondern Werkzeug der Profis in deinem
Alltag, meist unsichtbar im Hintergrund.

---

## Und wofür kannst DU Go später nutzen?

> *Einige reale Ideen von Junior-Entwicklerinnen:*

- Eine eigene Web-API (z. B. für deine Wetter-App)
- Ein CLI-Tool, das im Terminal Aufgaben erleichtert
- Einen Chatbot mit gleichzeitigen Verbindungen
- Ein Programm, das tausende Dateien verarbeitet
- Eine kleine Cloud-App auf AWS, GCP oder Azure
- Einen URL-Shortener mit eigener Datenbank
- Automatisierte Tests für andere Programme

All das geht mit Go, und alles davon lernst du in diesem Kurs.

---

## Dein Weg in den nächsten 90 Minuten

> *Vom "Hallo Welt" bis zur eigenen Funktion.*

| Zeit | Kapitel | Thema |
|---------|-----------|-------|
| 5 min | `00_intro` | Dein erstes Programm |
| 15 min | `01_basics` | Variablen und Texte |
| 35 min | `02_control_flow` | Entscheidungen und Schleifen |
| 60 min | `03_functions` | Wiederverwendbarer Code |
| 90 min | `04_data_structures` | Listen und Tabellen |

Alles weitere findest du im **[Lernpfad](../LEARNING_PATH.md)**.

---

## Bereit? Dann auf zu [`00_intro`](../00_intro/)!

> Kleiner Tipp: Du musst NICHTS installieren, um weiterzumachen.
> Öffne [go.dev/play](https://go.dev/play/), kopiere ein Code-Schnipsel
> aus `00_intro` und klicke **"Run"**.

Kapitel abgeschlossen.

---

## Quellcode

Das Go-Programm, das diese Geschichte als Terminal-Output erzeugt, liegt unter
[`00_geschichte/main.go`](main.go). Führe es aus mit:

```bash
go run 00_geschichte/main.go
```

---

## Übungen (alle optional, keinerlei Druck)

1. Welche der 12 Sprachen in der Tabelle hat dich am meisten überrascht? Recherchiere im Internet, was heute damit gebaut wird.
2. Suche "Hello World in Rust" und vergleiche: wieviel mehr oder weniger Code braucht es gegenüber Go?
3. Bonus: Stell dir vor, du müsstest 1.000.000 Anfragen pro Sekunde beantworten. Welche der genannten Sprachen wäre deiner Meinung nach die erste Wahl, und warum?
4. Bonus: Welche Firma aus der Liste oben beeindruckt dich am meisten, dass sie Go einsetzt? Lese 5 Minuten über ihr Go-Projekt auf deren Webseite.
