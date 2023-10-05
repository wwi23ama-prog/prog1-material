package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(1))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return Factorial(n-1) * n
}
