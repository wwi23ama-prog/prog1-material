package factorial

// Factorial5 berechnet die Fakultät der Zahl 5 in Einzelschritten.
func Factorial5SingleSteps() int {
	result := 1 // Startwert
	result = result * 2
	result = result * 3
	result = result * 4
	result = result * 5

	return result
}

// Factorial5Loop berechnet die Fakultät der Zahl 5 in einer Schleife.
func Factorial5Loop() int {
	result := 1 // Startwert
	for i := 2; i <= 5; i++ {
		result = result * i
	}

	return result
}

// FactorialNLoop berechnet die Fakultät einer beliebigen Zahl n.
func FactorialNLoop(n int) int {
	result := 1 // Startwert
	for i := 2; i <= n; i++ {
		result = result * i
	}

	return result
}

// FactorialNLoopBackwards berechnet die Fakultät einer beliebigen Zahl n.
// Die Schleife läuft rückwärts.
func FactorialNLoopBackwards(n int) int {
	result := 1 // Startwert
	for i := n; i >= 1; i-- {
		result = result * i
	}

	return result
}

// FactorialNRecursive berechnet die Fakultät einer beliebigen Zahl n.
// Die Funktion ruft sich selbst auf.
func FactorialNRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialNRecursive(n-1)
}
