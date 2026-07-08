// Kapitel 0: Geschichte der Programmiersprachen
//
// Lernziele:
// - Wann und warum Programmiersprachen erfunden wurden
// - Was die bekanntesten Sprachen ausmacht (Stärken, Schwächen, Use Cases)
// - WARUM es sich lohnt, heute Go zu lernen
//
// Herzlich willkommen!
// Keine Sorge, wenn du noch nie programmiert hast.
// Dieses Kapitel enthält KEINEN Code, den du schreiben musst.
//
// Reihenfolge für Anfänger: 00_erste_schritte (optional) → 00_geschichte →
// 00_intro → 01_basics → 02_control_flow → 03_functions → 04_data_structures → …
// Du kannst das Kapitel auch einfach im Browser im [Go Playground](https://go.dev/play/) lesen.

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Diese Konstanten machen den Code lesbarer.
const (
	trennlinie = "────────────────────────────────────────────────────────────"
	herz       = ""
	buch       = ""
	werkzeug   = ""
	goLogo     = ""
	praxis     = ""
	tipp       = ""
	runePfeil  = ""
	boxBreite  = 60
)

// typSprache bündelt die Daten jeder Programmiersprache.
type typSprache struct {
	name      string
	jahr      int
	kategorie string // Wissenschaft, Wirtschaft, Web, …
	stärke    string // eine Hauptstärke
	schwäche  string // eine bekannte Schwäche
	heute     string // wofür sie heute genutzt wird
	hello     string // ein typisches "Hello World"
}

// daten enthält alle Sprachen, die wir vorstellen.
var daten = []typSprache{
	{"FORTRAN", 1957, "Wissenschaft", "superschnell für numerische Berechnungen", "veraltete Syntax für moderne Augen", "Wetter- & Klimamodelle, Physik-Simulationen", "PRINT *, 'Hello'"},
	{"LISP", 1958, "Forschung", "elegante Mathematik, Code = Daten", "viele Klammern schrecken ab", "KI-Forschung, Sprachverarbeitung", "(print \"Hello\")"},
	{"COBOL", 1959, "Wirtschaft", "extrem gut für Buchhaltung und Reports", "sehr geschwätzig, schwer zu modernisieren", "Banken, Versicherungen, Behörden", "DISPLAY 'Hello'."},
	{"BASIC", 1964, "Bildung", "sehr einfach zu lernen", "untauglich für große Projekte", "Heimcomputer in den 80ern, heute Nische", "10 PRINT \"Hello\""},
	{"Pascal", 1970, "Lehre", "klar und didaktisch", "kaum noch in der Industrie", "Schulen, historisches Delphi", "writeln('Hello');"},
	{"C", 1972, "Systeme", "extrem schnell, vollständige Speicherkontrolle", "Fehler im Speicher sind berüchtigt", "Betriebssysteme, Datenbanken, Go selbst", "printf(\"Hello\\n\");"},
	{"C++", 1985, "Systeme", "C mit Objektorientierung und Templates", "viele Wege, es zu nutzen — viele Wege, Fehler zu machen", "Spiele, Browser, Bürosoftware", "cout << \"Hello\" << endl;"},
	{"Python", 1991, "Allzweck", "sehr lesbar, riesige Bibliotheken", "langsam, schwer für Mobile/Spiele", "Data Science, KI, Skripte", "print('Hello')"},
	{"Java", 1995, "Enterprise", "einmal schreiben, überall laufen", "viel Speicher und Boilerplate", "Große Firmen, Android-Apps", "System.out.println(\"Hello\");"},
	{"JavaScript", 1995, "Web", "läuft in jedem Browser, riesiges Ökosystem", "chaotisch gewachsen, viele Frameworks", "Web-Frontends, mit Node.js auch Server", "console.log('Hello');"},
	{"PHP", 1995, "Server-Web", "schnell zur lauffähigen Webseite", "unübersichtlicher Code bei großen Projekten", "WordPress, Wikipedia, große Teile des Webs", "echo 'Hello';"},
	{"Go", 2009, "Cloud & Server", "schnell wie C, einfach wie Python, sicher wie Java", "weniger Bibliotheken als Python/JS, keine Klassen-Vererbung", "Cloud, Microservices, CLI, DevOps", "fmt.Println(\"Hello\")"},
}

// zeigeSprache druckt die Detailkarte einer Sprache.
func zeigeSprache(s typSprache) {
	fmt.Printf(" %-9s (%d) – %s\n", s.name, s.jahr, s.kategorie)
	fmt.Printf(" So sieht Hallo Welt aus: %s\n", s.hello)
	fmt.Printf(" Stärke: %s\n", s.stärke)
	fmt.Printf(" Schwäche: %s\n", s.schwäche)
	fmt.Printf(" Heute: %s\n", s.heute)
	fmt.Println()
}

// rubrik druckt eine Sektion mit Überschrift und kurzer Erklärung.
func rubrik(titel string, erklaerung string) {
	fmt.Println()
	fmt.Println(trennlinie)
	fmt.Printf(" %s\n", titel)
	fmt.Printf(" %s\n", erklaerung)
	fmt.Println(trennlinie)
}

// codeOutputBlock druckt einen ASCII-Rahmen, der CODE- und OUTPUT-Bereich
// sauber voneinander trennt. Alle Längenberechnungen sind rune-basiert, damit
// Umlaute und Emojis die Box nicht auseinanderziehen.
func codeOutputBlock(titel string, codeLines []string, output string) {
	border := strings.Repeat("─", boxBreite)

	fmt.Println()
	fmt.Println(" ┌" + border + "┐")
	header := " " + titel + " "
	hr := utf8.RuneCountInString(header)
	if hr < boxBreite-2 {
		header = header + strings.Repeat("─", boxBreite-2-hr)
	} else if hr > boxBreite-2 {
		// Auf Rune-Grenze kürzen, nicht auf Byte-Grenze.
		runes := []rune(header)
		header = string(runes[:boxBreite-2])
	}
	fmt.Println(" │" + header + "│")
	fmt.Println(" │" + strings.Repeat("─", boxBreite) + "│")

	fmt.Println(" │ CODE (das, was der Entwickler tippt):" +
		strings.Repeat(" ", boxBreite-utf8.RuneCountInString(" │ CODE (das, was der Entwickler tippt):")) + "│")
	for _, line := range codeLines {
		padded := " │ " + line
		padded = padToBox(padded)
		fmt.Println(padded)
	}

	fmt.Println(" │" + strings.Repeat("─", boxBreite) + "│")
	fmt.Println(" │ OUTPUT (das, was der Computer zeigt):" +
		strings.Repeat(" ", boxBreite-utf8.RuneCountInString(" │ OUTPUT (das, was der Computer zeigt):")) + "│")
	padded := " │ " + output
	padded = padToBox(padded)
	fmt.Println(padded)
	fmt.Println(" └" + border + "┘")
}

// padToBox kürzt oder füllt eine bereits mit " │" beginnende Zeile
// bis zur konfigurierten Box-Breite auf und schließt mit │ ab.
// Rune-basiert, damit Umlaute/Emojis die Ausrichtung halten.
func padToBox(line string) string {
	const prefix = " │"
	const suffix = "│"
	content := strings.TrimPrefix(line, prefix)
	inner := boxBreite - 1
	cr := utf8.RuneCountInString(content)
	if cr > inner {
		// Auf Rune-Grenze kürzen + Ellipsis (1 Rune).
		runes := []rune(content)
		content = string(runes[:inner-1]) + "…"
		cr = utf8.RuneCountInString(content)
	}
	return prefix + content + strings.Repeat(" ", inner-cr) + suffix
}

func main() {
	// ============================================================
	// Begrüßung — einladend, kurz, ein Versprechen.
	// ============================================================
	fmt.Println()
	fmt.Printf("%s Schön, dass du da bist!\n", herz)
	fmt.Println()
	fmt.Println(" Dies ist kein Programmier-Test, sondern eine kleine")
	fmt.Println(" Geschichte — erzählt von Computern, Ideen und Menschen,")
	fmt.Println(" die vor dir dieselben Fragen gestellt haben.")
	fmt.Println()
	fmt.Println(" Dauer zum Lesen: ca. 5 Minuten.")
	fmt.Println(" Am besten mit einem Getränk in der Hand.")
	fmt.Println()

	// ============================================================
	// Teil 1 — Warum gibt es überhaupt Programmiersprachen?
	// ============================================================
	rubrik(
		"Wie kommt eine Programmiersprache in die Welt?",
		"Jede Sprache entsteht, weil jemand ein Problem hatte.",
	)
	fmt.Println()
	fmt.Println(" Stell dir vor: Computer verstehen nur Nullen und Einsen.")
	fmt.Println(" Maschinensprache sieht so aus: 01001000 01100101 01101100 …")
	fmt.Println()
	fmt.Println(" Schon nach kurzer Zeit fragten sich findige Köpfe:")
	fmt.Println(" \"Kann man Computern nicht einfach sagen, was sie tun")
	fmt.Println(" sollen — in Worten, die wir Menschen verstehen?\"")
	fmt.Println()

	// Vergleich: so sieht \"Hallo Welt\" in Go aus.
	codeOutputBlock(
		"So sieht \"Hallo Welt\" in Go aus",
		[]string{`fmt.Println("Hallo Welt!")`},
		"Hallo Welt!",
	)

	// ============================================================
	// Teil 2 — Die fünf Epochen im Überblick.
	// ============================================================
	rubrik(
		"Fünf Epochen, in denen sich die Welt verändert hat",
		"Von 1843 bis heute — eine Reise in fünf Etappen.",
	)

	fmt.Println()
	fmt.Println(" 1⃣ Die Anfänge (1843–1940)")
	fmt.Println(" Ideen ohne Computer — Berechnungspläne auf Papier.")
	fmt.Println()
	fmt.Println(" 2⃣ Erste Hochsprachen (1950–1969)")
	fmt.Println(" Maschinen werden leistungsfähiger —")
	fmt.Println(" Wörter ersetzen die Zahlenwüste.")
	fmt.Println()
	fmt.Println(" 3⃣ Strukturierte Programmierung (1970–1989)")
	fmt.Println(" \"Spaghetti-Code\" war ein echtes Problem.")
	fmt.Println(" Geordnete Strukturen und das legendäre C.")
	fmt.Println()
	fmt.Println(" 4⃣ Web-Ära (1990–2009)")
	fmt.Println(" Das Internet verändert alles.")
	fmt.Println(" Schnell schreiben ist wichtiger als schnell laufen.")
	fmt.Println()
	fmt.Println(" 5⃣ Cloud & Moderne (seit 2009)")
	fmt.Println(" Heute laufen Millionen Aufgaben gleichzeitig.")
	fmt.Println(" Neue Sprachen — Go, Rust, Kotlin — antworten darauf.")
	fmt.Println()

	// ============================================================
	// Teil 3 — Jede Sprache kurz und lebendig.
	// ============================================================
	rubrik(
		fmt.Sprintf("%s Die wichtigsten Sprachen — wofür sie gut sind", buch),
		"Du musst keine einzige davon heute lernen.",
	)
	fmt.Println()
	fmt.Println(" Diese Übersicht soll dir nur ein Gefühl dafür geben,")
	fmt.Println(" warum es so viele Sprachen gibt — und warum jede")
	fmt.Println(" ihre Daseinsberechtigung hat.")
	fmt.Println()

	for _, s := range daten {
		zeigeSprache(s)
	}

	// ============================================================
	// Teil 4 — Warum es nicht 'die beste Sprache' gibt.
	// ============================================================
	rubrik(
		"Es gibt keine \"beste Sprache\" — nur das beste Werkzeug",
		"Sprachen sind wie Werkzeuge in einer Werkstatt.",
	)
	fmt.Println()
	fmt.Println(" Ein Hammer ist nicht \"besser\" als eine Säge.")
	fmt.Println(" Es kommt darauf an, was du bauen willst.")
	fmt.Println()
	fmt.Println(" Wer eine App für 10.000 Menschen gleichzeitig schreibt,")
	fmt.Println(" trifft andere Entscheidungen als jemand, der eine")
	fmt.Println(" Webseite für einen kleinen Verein baut.")
	fmt.Println()
	fmt.Println(" Und genau deshalb gibt es heute über 200 Sprachen — und")
	fmt.Println(" jede Woche neue. Das ist kein Chaos, sondern Vielfalt.")
	fmt.Println()

	// ============================================================
	// Teil 5 — WARUM gerade Go?
	// ============================================================
	rubrik(
		fmt.Sprintf("%s Und warum schauen wir uns jetzt Go an?", goLogo),
		"Go wurde 2009 bei Google geboren — aus einem echten Problem.",
	)
	fmt.Println()
	fmt.Println(" Google hatte Anfang der 2000er ein gewaltiges Problem:")
	fmt.Println(" • Tausende Entwickler schreiben Code.")
	fmt.Println(" • Millionen Nutzer sprechen gleichzeitig mit den Servern.")
	fmt.Println(" • Der Code-Check dauerte bei großen Projekten 30+ Minuten.")
	fmt.Println(" • Die Computer verbrauchten viel zu viel Strom und Speicher.")
	fmt.Println()
	fmt.Println(" Drei erfahrene Entwickler machten sich daran, eine")
	fmt.Println(" Sprache zu entwerfen, die ALLE diese Probleme adressiert:")
	fmt.Println(" Rob Pike, Ken Thompson (Mit-Erfinder von Unix und C)")
	fmt.Println(" und Robert Griesemer (Schöpfer der Java-HotSpot-VM).")
	fmt.Println()
	fmt.Println(" Ihr Ziel war nicht \"die perfekte Sprache\", sondern:")
	fmt.Println(" die richtige Sprache für die Aufgaben von heute.")
	fmt.Println()

	// Vergleich: so kurz ist ein komplettes Go-Programm.
	codeOutputBlock(
		"So kurz ist ein komplettes Go-Programm",
		[]string{
			"package main",
			`import "fmt"`,
			`func main() { fmt.Println("Hallo, Geschichte!") }`,
		},
		"Hallo, Geschichte!",
	)

	// ============================================================
	// Teil 6 — Was Go so besonders macht.
	// ============================================================
	rubrik(
		fmt.Sprintf("%s Was Go anders macht als andere Sprachen", werkzeug),
		"Sechs Eigenschaften, die du beim Lernen wiedererkennen wirst.",
	)
	fmt.Println()
	fmt.Println(" 1. Einfach zu lesen, einfach zu lernen")
	fmt.Println(" Go hat nur 25 Schlüsselwörter. C hat über 30.")
	fmt.Println(" Wer Go liest, versteht es — auch wenn er es nicht selbst schreibt.")
	fmt.Println()
	fmt.Println(" 2. Superschnelle Kompilierung")
	fmt.Println(" Große Projekte wie Kubernetes (~3 Mio. Zeilen) bauen in unter 60 Sek.")
	fmt.Println(" Vergleich: C++-Projekte ähnlicher Größe brauchen Minuten.")
	fmt.Println()
	fmt.Println(" 3. Eingebaute Nebenläufigkeit (Concurrency)")
	fmt.Println(" Tausende \"Goroutinen\" starten in einer einzigen Datei.")
	fmt.Println(" Wer 10.000 Aufgaben gleichzeitig erledigen muss:")
	fmt.Println(" Go ist oft die einfachste Sprache dafür.")
	fmt.Println()
	fmt.Println(" 4. Eine Datei, ein Programm")
	fmt.Println(" Go baut eine einzelne .exe-Datei (oder Linux-Binary).")
	fmt.Println(" Kein Runtime, keine VM, keine Frameworks-Liste. Starten, fertig.")
	fmt.Println()
	fmt.Println(" 5. Speicher wird automatisch aufgeräumt (Garbage Collection)")
	fmt.Println(" In C/C++ muss man Speicher selbst freigeben und vergisst es leicht.")
	fmt.Println(" In Java kostet das viel Performance. Go macht beides besser.")
	fmt.Println()
	fmt.Println(" 6. Eine Sprache, ein Stil")
	fmt.Println(" Es gibt nur einen verbreiteten Weg, Code zu formatieren.")
	fmt.Println(" \"go fmt\" formatiert deinen Code automatisch korrekt.")
	fmt.Println(" Streit über Einrückung? Gibt es in Go nicht. ")
	fmt.Println()

	// ============================================================
	// Teil 7 — Wo Go heute überall läuft.
	// ============================================================
	rubrik(
		fmt.Sprintf("%s Wo du Go heute überall im Einsatz siehst", praxis),
		"Auch wenn du es nicht merkst.",
	)
	fmt.Println()
	fmt.Println(" Berühmte Open-Source-Projekte, die in Go geschrieben sind:")
	fmt.Println(" Docker – Container, die deine Apps starten")
	fmt.Println(" Kubernetes – Verwaltet tausende Container")
	fmt.Println(" Terraform – Beschreibt Cloud-Infrastruktur als Code")
	fmt.Println(" Hugo – Diese statischen Webseiten")
	fmt.Println(" Caddy – Webserver mit automatischem HTTPS")
	fmt.Println(" Prometheus – Monitoring für moderne Systeme")
	fmt.Println(" CockroachDB – Verteilte SQL-Datenbank")
	fmt.Println()
	fmt.Println(" Große Firmen, die Go produktiv einsetzen:")
	fmt.Println(" Google – Ursprung und viele interne Tools")
	fmt.Println(" Uber – Geokoordinaten und Matching in Echtzeit")
	fmt.Println(" Twitch – Video-Pipeline für Live-Streams")
	fmt.Println(" Stripe – Teile der Zahlungs-Infrastruktur")
	fmt.Println(" Cloudflare – DNS, Edge-Computing, Sicherheit")
	fmt.Println(" Dropbox – Magischer Datei-Sync")
	fmt.Println()
	fmt.Println(" Go ist keine \"Geheimtip\"-Sprache mehr, sondern Werkzeug")
	fmt.Println(" der Profis in deinem Alltag — meist unsichtbar im Hintergrund.")
	fmt.Println()

	// ============================================================
	// Teil 8 — Wofür du Go in Zukunft nutzen kannst.
	// ============================================================
	rubrik(
		"Und wofür kannst DU Go später nutzen?",
		"Einige reale Ideen von Junior-Entwicklerinnen:",
	)
	fmt.Println()
	fmt.Println(" Eine eigene Web-API (z. B. für deine Wetter-App)")
	fmt.Println(" Ein CLI-Tool, das im Terminal Aufgaben erleichtert")
	fmt.Println(" Einen Chatbot mit gleichzeitigen Verbindungen")
	fmt.Println(" Ein Programm, das tausende Dateien verarbeitet")
	fmt.Println(" Eine kleine Cloud-App auf AWS, GCP oder Azure")
	fmt.Println(" Einen URL-Shortener mit eigener Datenbank")
	fmt.Println(" Automatisierte Tests für andere Programme")
	fmt.Println()
	fmt.Println(" All das geht mit Go — und alles davon lernst du in diesem Kurs.")
	fmt.Println()

	// ============================================================
	// Teil 9 — Was dich in den nächsten Stunden erwartet.
	// ============================================================
	rubrik(
		"Dein Weg in den nächsten 90 Minuten",
		"Vom \"Hallo Welt\" bis zur eigenen Funktion.",
	)
	fmt.Println()
	fmt.Printf(" 5 min %s Kapitel 00_intro: dein erstes Programm\n", runePfeil)
	fmt.Printf(" 15 min %s Kapitel 01_basics: Variablen und Texte\n", runePfeil)
	fmt.Printf(" 35 min %s Kapitel 02_control: Entscheidungen & Schleifen\n", runePfeil)
	fmt.Printf(" 60 min %s Kapitel 03_functions: wiederverwendbarer Code\n", runePfeil)
	fmt.Printf(" 90 min %s Kapitel 04_data: Listen & Tabellen\n", runePfeil)
	fmt.Println()
	fmt.Println(" Alles weitere findest du im [Lernpfad](../LEARNING_PATH.md).")
	fmt.Println()

	// ============================================================
	// Abschluss
	// ============================================================
	fmt.Println()
	fmt.Println(trennlinie)
	fmt.Printf(" %s Bereit? Dann auf zu [00_intro](../00_intro/)!\n", herz)
	fmt.Println(trennlinie)
	fmt.Println()
	fmt.Printf(" %s Kleiner Tipp: Du musst NICHTS installieren, um weiterzumachen.\n", tipp)
	fmt.Println(" Öffne [go.dev/play](https://go.dev/play/), kopiere ein")
	fmt.Println(" Code-Schnipsel aus 00_intro und klicke \"Run\".")
	fmt.Println()
	fmt.Println(" Kapitel abgeschlossen!")
}

// -------------------------------------------------------------------
// Übungen (alle optional, keinerlei Druck)
//
// 1. Welche der 12 Sprachen in der Tabelle hat dich am meisten
// überrascht? Recherchiere im Internet, was heute damit gebaut wird.
// 2. Suche "Hello World in Rust" und vergleiche: wieviel mehr oder
// weniger Code braucht es gegenüber Go?
// 3. Bonus: Stell dir vor, du müsstest 1.000.000 Anfragen pro Sekunde
// beantworten. Welche der genannten Sprachen wäre deiner Meinung
// nach die erste Wahl — und warum?
// 4. Bonus: Welche Firma aus der Liste oben beeindruckt dich am
// meisten, dass sie Go einsetzt? Lese 5 Minuten über ihr
// Go-Projekt auf deren Webseite.
// -------------------------------------------------------------------
