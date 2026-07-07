# 📖 Geschichte der Programmiersprachen

> **Hinweis:** Dieses Dokument wurde automatisch aus `00_geschichte/main.go` generiert. Es enthält den vollständigen Output des Go-Programms — Theorie, ASCII-Boxen mit Code-Beispielen, Programmiersprachen-Tabelle und Übungen.

---

    
    💛  Schön, dass du da bist!
    
       Dies ist kein Programmier-Test, sondern eine kleine
       Geschichte — erzählt von Computern, Ideen und Menschen,
       die vor dir dieselben Fragen gestellt haben.
    
       ⏱️  Dauer zum Lesen: ca. 5 Minuten.
       ☕  Am besten mit einem Getränk in der Hand.
    
    
    ────────────────────────────────────────────────────────────
      Wie kommt eine Programmiersprache in die Welt?
      Jede Sprache entsteht, weil jemand ein Problem hatte.
    ────────────────────────────────────────────────────────────
    
       Stell dir vor: Computer verstehen nur Nullen und Einsen.
       Maschinensprache sieht so aus: 01001000 01100101 01101100 …
    
       Schon nach kurzer Zeit fragten sich findige Köpfe:
       "Kann man Computern nicht einfach sagen, was sie tun
       sollen — in Worten, die wir Menschen verstehen?"
    
    
      ┌────────────────────────────────────────────────────────────┐
      │ So sieht "Hallo Welt" in Go aus ─────────────────────────│
      │────────────────────────────────────────────────────────────│
      │  CODE  (das, was der Entwickler tippt):                 │
      │   fmt.Println("Hallo Welt!")                              │
      │────────────────────────────────────────────────────────────│
      │  OUTPUT  (das, was der Computer zeigt):                 │
      │   Hallo Welt!                                             │
      └────────────────────────────────────────────────────────────┘
    
    ────────────────────────────────────────────────────────────
      Fünf Epochen, in denen sich die Welt verändert hat
      Von 1843 bis heute — eine Reise in fünf Etappen.
    ────────────────────────────────────────────────────────────
    
       1️⃣  Die Anfänge (1843–1940)
           Ideen ohne Computer — Berechnungspläne auf Papier.
    
       2️⃣  Erste Hochsprachen (1950–1969)
           Maschinen werden leistungsfähiger —
           Wörter ersetzen die Zahlenwüste.
    
       3️⃣  Strukturierte Programmierung (1970–1989)
           "Spaghetti-Code" war ein echtes Problem.
           Geordnete Strukturen und das legendäre C.
    
       4️⃣  Web-Ära (1990–2009)
           Das Internet verändert alles.
           Schnell schreiben ist wichtiger als schnell laufen.
    
       5️⃣  Cloud & Moderne (seit 2009)
           Heute laufen Millionen Aufgaben gleichzeitig.
           Neue Sprachen — Go, Rust, Kotlin — antworten darauf.
    
    
    ────────────────────────────────────────────────────────────
      📖 Die wichtigsten Sprachen — wofür sie gut sind
      Du musst keine einzige davon heute lernen.
    ────────────────────────────────────────────────────────────
    
       Diese Übersicht soll dir nur ein Gefühl dafür geben,
       warum es so viele Sprachen gibt — und warum jede
       ihre Daseinsberechtigung hat.
    
       ★ FORTRAN   (1957) – Wissenschaft
           So sieht Hallo Welt aus:  PRINT *, 'Hello'
           Stärke:    superschnell für numerische Berechnungen
           Schwäche:  veraltete Syntax für moderne Augen
           Heute:     Wetter- & Klimamodelle, Physik-Simulationen
    
       ★ LISP      (1958) – Forschung
           So sieht Hallo Welt aus:  (print "Hello")
           Stärke:    elegante Mathematik, Code = Daten
           Schwäche:  viele Klammern schrecken ab
           Heute:     KI-Forschung, Sprachverarbeitung
    
       ★ COBOL     (1959) – Wirtschaft
           So sieht Hallo Welt aus:  DISPLAY 'Hello'.
           Stärke:    extrem gut für Buchhaltung und Reports
           Schwäche:  sehr geschwätzig, schwer zu modernisieren
           Heute:     Banken, Versicherungen, Behörden
    
       ★ BASIC     (1964) – Bildung
           So sieht Hallo Welt aus:  10 PRINT "Hello"
           Stärke:    sehr einfach zu lernen
           Schwäche:  untauglich für große Projekte
           Heute:     Heimcomputer in den 80ern, heute Nische
    
       ★ Pascal    (1970) – Lehre
           So sieht Hallo Welt aus:  writeln('Hello');
           Stärke:    klar und didaktisch
           Schwäche:  kaum noch in der Industrie
           Heute:     Schulen, historisches Delphi
    
       ★ C         (1972) – Systeme
           So sieht Hallo Welt aus:  printf("Hello\n");
           Stärke:    extrem schnell, vollständige Speicherkontrolle
           Schwäche:  Fehler im Speicher sind berüchtigt
           Heute:     Betriebssysteme, Datenbanken, Go selbst
    
       ★ C++       (1985) – Systeme
           So sieht Hallo Welt aus:  cout << "Hello" << endl;
           Stärke:    C mit Objektorientierung und Templates
           Schwäche:  viele Wege, es zu nutzen — viele Wege, Fehler zu machen
           Heute:     Spiele, Browser, Bürosoftware
    
       ★ Python    (1991) – Allzweck
           So sieht Hallo Welt aus:  print('Hello')
           Stärke:    sehr lesbar, riesige Bibliotheken
           Schwäche:  langsam, schwer für Mobile/Spiele
           Heute:     Data Science, KI, Skripte
    
       ★ Java      (1995) – Enterprise
           So sieht Hallo Welt aus:  System.out.println("Hello");
           Stärke:    einmal schreiben, überall laufen
           Schwäche:  viel Speicher und Boilerplate
           Heute:     Große Firmen, Android-Apps
    
       ★ JavaScript (1995) – Web
           So sieht Hallo Welt aus:  console.log('Hello');
           Stärke:    läuft in jedem Browser, riesiges Ökosystem
           Schwäche:  chaotisch gewachsen, viele Frameworks
           Heute:     Web-Frontends, mit Node.js auch Server
    
       ★ PHP       (1995) – Server-Web
           So sieht Hallo Welt aus:  echo 'Hello';
           Stärke:    schnell zur lauffähigen Webseite
           Schwäche:  unübersichtlicher Code bei großen Projekten
           Heute:     WordPress, Wikipedia, große Teile des Webs
    
       ★ Go        (2009) – Cloud & Server
           So sieht Hallo Welt aus:  fmt.Println("Hello")
           Stärke:    schnell wie C, einfach wie Python, sicher wie Java
           Schwäche:  weniger Bibliotheken als Python/JS, keine Klassen-Vererbung
           Heute:     Cloud, Microservices, CLI, DevOps
    
    
    ────────────────────────────────────────────────────────────
      Es gibt keine "beste Sprache" — nur das beste Werkzeug
      Sprachen sind wie Werkzeuge in einer Werkstatt.
    ────────────────────────────────────────────────────────────
    
       Ein Hammer ist nicht "besser" als eine Säge.
       Es kommt darauf an, was du bauen willst.
    
       Wer eine App für 10.000 Menschen gleichzeitig schreibt,
       trifft andere Entscheidungen als jemand, der eine
       Webseite für einen kleinen Verein baut.
    
       Und genau deshalb gibt es heute über 200 Sprachen — und
       jede Woche neue. Das ist kein Chaos, sondern Vielfalt.
    
    
    ────────────────────────────────────────────────────────────
      🐹 Und warum schauen wir uns jetzt Go an?
      Go wurde 2009 bei Google geboren — aus einem echten Problem.
    ────────────────────────────────────────────────────────────
    
       Google hatte Anfang der 2000er ein gewaltiges Problem:
       • Tausende Entwickler schreiben Code.
       • Millionen Nutzer sprechen gleichzeitig mit den Servern.
       • Der Code-Check dauerte bei großen Projekten 30+ Minuten.
       • Die Computer verbrauchten viel zu viel Strom und Speicher.
    
       Drei erfahrene Entwickler machten sich daran, eine
       Sprache zu entwerfen, die ALLE diese Probleme adressiert:
       Rob Pike, Ken Thompson (Mit-Erfinder von Unix und C)
       und Robert Griesemer (Schöpfer der Java-HotSpot-VM).
    
       Ihr Ziel war nicht "die perfekte Sprache", sondern:
       die richtige Sprache für die Aufgaben von heute.
    
    
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
    
    ────────────────────────────────────────────────────────────
      🛠️ Was Go anders macht als andere Sprachen
      Sechs Eigenschaften, die du beim Lernen wiedererkennen wirst.
    ────────────────────────────────────────────────────────────
    
       1. Einfach zu lesen, einfach zu lernen
          Go hat nur 25 Schlüsselwörter. C hat über 30.
          Wer Go liest, versteht es — auch wenn er es nicht selbst schreibt.
    
       2. Superschnelle Kompilierung
          Große Projekte wie Kubernetes (~3 Mio. Zeilen) bauen in unter 60 Sek.
          Vergleich: C++-Projekte ähnlicher Größe brauchen Minuten.
    
       3. Eingebaute Nebenläufigkeit (Concurrency)
          Tausende "Goroutinen" starten in einer einzigen Datei.
          Wer 10.000 Aufgaben gleichzeitig erledigen muss:
          Go ist oft die einfachste Sprache dafür.
    
       4. Eine Datei, ein Programm
          Go baut eine einzelne .exe-Datei (oder Linux-Binary).
          Kein Runtime, keine VM, keine Frameworks-Liste. Starten, fertig.
    
       5. Speicher wird automatisch aufgeräumt (Garbage Collection)
          In C/C++ muss man Speicher selbst freigeben und vergisst es leicht.
          In Java kostet das viel Performance. Go macht beides besser.
    
       6. Eine Sprache, ein Stil
          Es gibt nur einen verbreiteten Weg, Code zu formatieren.
          "go fmt" formatiert deinen Code automatisch korrekt.
          Streit über Einrückung? Gibt es in Go nicht. 🎉
    
    
    ────────────────────────────────────────────────────────────
      ✨ Wo du Go heute überall im Einsatz siehst
      Auch wenn du es nicht merkst.
    ────────────────────────────────────────────────────────────
    
       Berühmte Open-Source-Projekte, die in Go geschrieben sind:
       🐳  Docker             – Container, die deine Apps starten
       ☸️   Kubernetes         – Verwaltet tausende Container
       🏗️  Terraform          – Beschreibt Cloud-Infrastruktur als Code
       📝  Hugo               – Diese statischen Webseiten
       🌐  Caddy              – Webserver mit automatischem HTTPS
       📊  Prometheus         – Monitoring für moderne Systeme
       🗄️  CockroachDB        – Verteilte SQL-Datenbank
    
       Große Firmen, die Go produktiv einsetzen:
       🔍  Google             – Ursprung und viele interne Tools
       🚖  Uber               – Geokoordinaten und Matching in Echtzeit
       📺  Twitch             – Video-Pipeline für Live-Streams
       💳  Stripe             – Teile der Zahlungs-Infrastruktur
       ☁️  Cloudflare         – DNS, Edge-Computing, Sicherheit
       📦  Dropbox            – Magischer Datei-Sync
    
       Go ist keine "Geheimtip"-Sprache mehr, sondern Werkzeug
       der Profis in deinem Alltag — meist unsichtbar im Hintergrund.
    
    
    ────────────────────────────────────────────────────────────
      Und wofür kannst DU Go später nutzen?
      Einige reale Ideen von Junior-Entwicklerinnen:
    ────────────────────────────────────────────────────────────
    
       🌐  Eine eigene Web-API (z. B. für deine Wetter-App)
       🛠️  Ein CLI-Tool, das im Terminal Aufgaben erleichtert
       🤖  Einen Chatbot mit gleichzeitigen Verbindungen
       📁  Ein Programm, das tausende Dateien verarbeitet
       ☁️  Eine kleine Cloud-App auf AWS, GCP oder Azure
       🔎  Einen URL-Shortener mit eigener Datenbank
       🧪  Automatisierte Tests für andere Programme
    
       All das geht mit Go — und alles davon lernst du in diesem Kurs.
    
    
    ────────────────────────────────────────────────────────────
      Dein Weg in den nächsten 90 Minuten
      Vom "Hallo Welt" bis zur eigenen Funktion.
    ────────────────────────────────────────────────────────────
    
       ⏱️   5 min   ➜  Kapitel 00_intro:    dein erstes Programm
       ⏱️  15 min   ➜  Kapitel 01_basics:   Variablen und Texte
       ⏱️  35 min   ➜  Kapitel 02_control:  Entscheidungen & Schleifen
       ⏱️  60 min   ➜  Kapitel 03_functions: wiederverwendbarer Code
       ⏱️  90 min   ➜  Kapitel 04_data:      Listen & Tabellen
    
       Alles weitere findest du im [Lernpfad](../LEARNING_PATH.md).
    
    
    ────────────────────────────────────────────────────────────
      💛  Bereit? Dann auf zu [00_intro](../00_intro/)!
    ────────────────────────────────────────────────────────────
    
       💡  Kleiner Tipp: Du musst NICHTS installieren, um weiterzumachen.
          Öffne [go.dev/play](https://go.dev/play/), kopiere ein
          Code-Schnipsel aus 00_intro und klicke "Run".
    
    ✅ Kapitel abgeschlossen!

---

## 📖 Quellcode

Das Go-Programm, das diesen Output erzeugt hat, liegt unter `00_geschichte/main.go`. Führe es aus mit:

```bash
go run 00_geschichte/main.go
```

Oder lies den Source direkt in [00_geschichte/main.go](main.go).
