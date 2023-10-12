package loops

import "fmt"

func ExampleListNumbers() {
	ListNumbers(5)
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleListNumbersBackwards() {
	ListNumbersBackwards(5)
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleListEvenNumbers() {
	ListEvenNumbers(5)
	// Output:
	// 0
	// 2
	// 4
}

func ExampleListMultiplesOf() {
	ListMultiplesOf(3, 10)
	// Output:
	// 0
	// 3
	// 6
	// 9
}

func ExampleListMultiplesOfBigSteps() {
	ListMultiplesOfBigSteps(3, 10)
	// Output:
	// 0
	// 3
	// 6
	// 9
}

func ExampleSumN() {
	fmt.Println(SumN(5))
	// Output:
	// 15
}

func ExampleSumNRecursive() {
	fmt.Println(SumNRecursive(5))
	// Output:
	// 15
}

func ExampleIsPrime() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(8))
	fmt.Println(IsPrime(9))
	fmt.Println(IsPrime(10))
	fmt.Println(IsPrime(11))
	fmt.Println(IsPrime(12))
	fmt.Println(IsPrime(13))
	fmt.Println(IsPrime(14))
	fmt.Println(IsPrime(15))
	fmt.Println(IsPrime(16))
	fmt.Println(IsPrime(17))
	fmt.Println(IsPrime(18))
	fmt.Println(IsPrime(19))
	fmt.Println(IsPrime(20))
	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// false
}

func ExampleSumWhileN() {
	fmt.Println(SumWhileN(5))
	// Output:
	// 15
}
