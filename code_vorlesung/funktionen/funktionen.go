package funktionen

import "fmt"

func Foo() {
	x := 3
	y := Bar(x) + 2*3
	fmt.Println(x + y)
}

func Bar(x int) int {
	x = x*3 + 4
	return x / 2
}
