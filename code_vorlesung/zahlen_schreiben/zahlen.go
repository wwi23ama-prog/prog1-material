package zahlenschreiben

import "golang.org/x/exp/slices"

// Digits erwartet eine Zahl und liefert eine Liste
// mit den Ziffern dieser Zahl.
func Digits(n int) []int {
	result := []int{}

	for n != 0 { // Solange n nicht 0 ist.
		// Letzte Ziffer bestimmen.
		lastdigit := n % 10 // "n modulo 10"
		result = append(result, lastdigit)

		// Letzte Ziffer von n abschneiden.
		n = n / 10
	}
	slices.Reverse(result)

	return result
}

// DigitStrings erwartet eine Liste von Ziffern und liefert
// eine Liste mit den entsprechenden Wörtern zurück.
func DigitStrings(d []int) []string {
	result := []string{}
	// TODO
	return result
}

// Compose erwartet eine Liste von Ziffer-Strings,
// fügt sie zusammen und baut notwendige Füllwörter ein
// (z.B. "und", "hundert" etc.).
func Compose(dstrings []string) string {
	result := ""
	// TODO
	return result
}

// Erwartet eine Zahl und liefert eine menschenlesbare String-Darstellung.
func IntToString(n int) string {
	// TODO
	return ""
}
