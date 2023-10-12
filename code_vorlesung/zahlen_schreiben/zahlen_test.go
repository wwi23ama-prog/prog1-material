package zahlenschreiben

import "fmt"

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
