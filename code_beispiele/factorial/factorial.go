package factorial

// Factorial berechnet die Fakultät einer Zahl.
func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
