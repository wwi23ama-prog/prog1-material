package rekursion

import "fmt"

func ExampleCountA() {
	fmt.Println(CountA_v3("agtRAvkeAx38"))

	// Output:
	// 2
}

func ExampleCountA_v4() {
	fmt.Println(CountA_v4("agtRAvkeAx38", 0))

	// Output:
	// 2
}

func ExampleCountA_v5() {
	fmt.Println(CountA_v5("agtRAvkeAx38", 0, 0))

	// Output:
	// 2
}

func ExampleCountA_v6() {
	fmt.Println(CountA_v7("agtRAvkeAx38", 0))

	// Output:
	// 2
}
