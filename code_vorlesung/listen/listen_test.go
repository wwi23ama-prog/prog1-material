package listen

import "fmt"

func ExampleIntLists() {
	IntLists()

	// Output:
}

func ExampleMakeListAscending() {
	x := 3
	l1 := MakeListAscending(x)
	fmt.Println(l1)

	// Output:
}

func ExamplePrettyPrint() {
	l1 := []int{5, 4, 3, 2, 1}
	PrettyPrint(l1, "Liste 1")

	l2 := []int{42, 38}
	PrettyPrint(l2, "Foo")

	// Output:
	// Liste 1: 5, 4, 3, 2, 1
	// Foo: 42, 38

}

func ExamplePrettyPrint2() {
	l1 := []int{5, 4, 3, 2, 1}
	PrettyPrint2(l1, "Liste 1")

	l2 := []int{42, 38}
	PrettyPrint2(l2, "Foo")

	// Output:
	// Liste 1: 5, 4, 3, 2, 1
	// Foo: 42, 38

}
