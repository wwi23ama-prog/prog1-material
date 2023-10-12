package zahlenschreiben

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(143))

	// Output:
	// [1 4 3]
}

func ExampleIntToString() {
	fmt.Println(IntToString(143))
	e1 := IntToString(58)
	fmt.Println(e1)
	fmt.Println(IntToString(7))

	// Output:
	// einhundertdreiundvierzig
	// achtundfünfzig
	// sieben
}
