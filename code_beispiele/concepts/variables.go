package concepts

import "fmt"

func IntVariables() {
	var n int // Variablendeklaration
	n = 42    // Variablenzuweisung
	k := 23   // Kurzschreibweise für Deklaration und Zuweisung

	fmt.Println(n, k, n+k)
}

func StringVariables() {
	s := "Hallo"
	t := "Welt"

	st := s + " " + t // Verkettung der Strings

	fmt.Println(st)
}

func ListVariables() {
	var l []int // leere Liste
	l = append(l, 10, 20, 30, 40, 50)

	fmt.Println(l)      // komplett ausgeben
	fmt.Println(l[1])   // Zweites Element ausgeben
	fmt.Println(l[1:3]) // Teil-Liste ausgeben
	l[1] = 42           // Wert ändern

	fmt.Println(l)
}
