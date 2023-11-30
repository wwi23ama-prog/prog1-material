package structs

// Datentyp für die Elemente eines Wörterbuchs.
// Präziser: Ein Deutsch-English-Wörterbuch.
// Einträge bestehen aus:
// - einem deutschen Wort
// - einem englischen Wort
type Entry struct {
	de string
	en string
}

// Welche Methoden brauchen die Elemente?
// - Übersetzen

func (e Entry) Translate() string {
	return e.en
}

// Datentyp für das Wörterbuch.
// Liste von Einträgen.
type Dictionary struct {
	entries []Entry
}

// Welche Methoden braucht das Wörterbuch?
// Add: fügt ein Element hinzu.
// Lookup: Liefert das englische Wort
//         zu einem deutschen Wort.
// (Sort: Sortiert die Einträge.)
//
// Konstruktor: Funktion, die ein neues
//              Dictionary liefert.

// Add fügt ein neues Element hinzu.
func (d *Dictionary) Add(de, en string) {
	entry := Entry{de, en}
	d.entries = append(d.entries, entry)
}

// Lookup
