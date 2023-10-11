package formen

import "fmt"

// PrintRow gibt eine Zeile mit dem angegebenen Zeichen aus, ohne dabei Zeilenumbrüche
// oder anderes Drumherum zu machen.
func PrintRow(n int, c string) {
	for i := 0; i < n; i++ {
		fmt.Print(c)
	}
}

// PrintRowStar gibt eine Zeile aus Sternen mit anschließendem Zeilenumbruch aus.
func PrintRowStar(n int) {
	PrintRow(n, "*")
	fmt.Println()
}

// PrintRowSpace gibt eine Zeile aus Leerzeichen aus, die mit Sternen begrenzt ist und
// einen Zeilenumbruch hat.
func PrintRowSpace(n int) {
	fmt.Print("*")
	PrintRow(n-2, " ")
	fmt.Println("*")
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	for i := 0; i < n; i++ {
		PrintRowStar(n)
	}
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	PrintRowStar(n)
	for i := 0; i < n-2; i++ {
		PrintRowSpace(n)
	}
	PrintRowStar(n)

	// Bessere Lösung, sobald PrintCustomSquare implementiert ist.
	// PrintCustomSquare(n, "*", " ")
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// TODO
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// TODO
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// TODO
}
