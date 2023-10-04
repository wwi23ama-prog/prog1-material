package factorial

import "fmt"

func ExampleFactorial5SingleSteps() {
	fmt.Println(Factorial5SingleSteps())
	// Output:
	// 120
}

func ExampleFactorial5Loop() {
	fmt.Println(Factorial5Loop())
	// Output:
	// 120
}

func ExampleFactorialNLoop() {
	f5 := FactorialNLoop(5)
	f3 := FactorialNLoop(3)
	f1 := FactorialNLoop(1)
	f0 := FactorialNLoop(0)

	fmt.Println(f5)
	fmt.Println(f3)
	fmt.Println(f1)
	fmt.Println(f0)

	// Output:
	// 120
	// 6
	// 1
	// 1
}

func ExampleFactorialNLoopBackwards() {
	fmt.Println(FactorialNLoopBackwards(5))
	fmt.Println(FactorialNLoopBackwards(3))
	fmt.Println(FactorialNLoopBackwards(1))
	fmt.Println(FactorialNLoopBackwards(0))

	// Output:
	// 120
	// 6
	// 1
	// 1
}
