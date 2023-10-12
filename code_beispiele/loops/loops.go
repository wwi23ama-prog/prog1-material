package loops

import "fmt"

/* Generelles Schema einer Schleife:
for <Start>; <Bedingung>; <Schritt> {
	// Schleifenkörper
}
*/

// ListNumbers ist eine einfache for-Schleife, die von 0 bis n-1 zählt und die Zahlen ausgibt.
func ListNumbers(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

// ListNumbersBackwards ist eine einfache for-Schleife, die von n bis 1 zählt und die Zahlen ausgibt.
func ListNumbersBackwards(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

// ListEvenNumbers gibt alle geraden Zahlen von 0 bis n-1 aus.
func ListEvenNumbers(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// ListMultiplesOf gibt alle Vielfachen von m von 0 bis n-1 aus.
func ListMultiplesOf(m, n int) {
	for i := 0; i < n; i++ {
		if i%m == 0 {
			fmt.Println(i)
		}
	}
}

// ListMultiplesOfBigSteps macht das Gleiche wie ListMultiplesOf,
// aber die Schleife läuft in m-Schritten.
func ListMultiplesOfBigSteps(m, n int) {
	for i := 0; i < n; i += m {
		fmt.Println(i)
	}
}

// SumN berechnet die Summe der Zahlen von 1 bis n.
func SumN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// SumNRecursive berechnet die Summe der Zahlen von 1 bis n.
// Die Funktion ruft sich selbst auf.
func SumNRecursive(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumNRecursive(n-1)
}

// IsPrime prüft, ob eine Zahl eine Primzahl ist.
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// SumWhileN berechnet die Summe der Zahlen von 1 bis n.
// Die Funktion verwendet eine while-Schleife.
func SumWhileN(n int) int {
	sum, i := 0, 1
	for i <= n {
		sum += i
		i++
	}
	return sum
}
