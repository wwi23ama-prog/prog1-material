package listen

import "fmt"

func IntLists() {
	// Array für 5 int-Werte, automatisch initialisiert mit 0.
	a := [5]int{}
	fmt.Println(a)

	// Array für 7 Werte, selbst initialisiert.
	b := [7]int{3, 4, 5, 6, 7, 8, 38}
	fmt.Println(b)
	b[0] = 42
	fmt.Println(b)

	// Leere Slice definieren.
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)

	// Elemente an s1 anhängen
	s1 = append(s1, 42)
	s1 = append(s1, 25, 75, 38)
	fmt.Println(s1)

	// Teil-Slice definieren
	s2 := s1[2:6]
	fmt.Println(s2)

	// Element an Stelle 3 aus s1 löschen
	s1 = append(s1[:3], s1[4:]...)
	//	s1 = append(s1[0:3], s1[4:len(s1)]...)
	fmt.Println(s1)

	s2[0] = 99999
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 555)
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 777, 777, 777)
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 888
	fmt.Println(s1)
	fmt.Println(s2)
}

// MakeListAscending erzeugt eine Liste der Länge n
// und füllt sie mit den Zahlen von 1 bis n auf.
func MakeListAscending(n int) []int {
	result := []int{}

	for i := 1; i <= n; i++ {
		result = append(result, i)
	}

	return result
}

// PrettyPrint erwartet eine Liste und einen Namen für die Liste.
// Die Funktion gibt die Liste in menschenlesbarer Form
// mit ihrem Namen aus.
