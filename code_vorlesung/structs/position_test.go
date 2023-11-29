package structs

import "fmt"

func ExamplePosition() {
	var p1 Position

	fmt.Println(p1)

	p1.x = 42

	fmt.Println(p1)

	p1.y = 38

	fmt.Println(p1)

	fmt.Println(p1.x)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p1.x)
	fmt.Printf("%T\n", p1.y)

	// Output:
}

func ExamplePosition_create() {
	// Deklaration mit var
	var p1 Position

	// Definition mit Angabe von Werten
	p2 := Position{1, 2}

	// Definition mit Auswahl der Attribute
	p3 := Position{y: 1, x: 42}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	// Output:
}
