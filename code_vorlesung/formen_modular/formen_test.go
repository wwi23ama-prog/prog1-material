package formen

import "fmt"

// Zusätzlicher Test (nicht in Vorlesung geschrieben, aber für eine Funktion aus der Vorlesung)
func ExamplePrintRow() {
	PrintRow(5, "A")
	fmt.Println()

	PrintRow(4, "B")
	fmt.Println()

	PrintRow(3, "C")
	fmt.Println()

	// Output:
	// AAAAA
	// BBBB
	// CCC
}

func ExamplePrintRowStar() {
	PrintRowStar(5)
	PrintRowStar(4)
	PrintRowStar(3)
	PrintRowStar(2)
	PrintRowStar(1)

	// Output:
	// *****
	// ****
	// ***
	// **
	// *
}

// Zusätzlicher Test (nicht in Vorlesung geschrieben, aber für eine Funktion aus der Vorlesung)
func ExamplePrintRowSpace() {
	PrintRowSpace(5)
	PrintRowSpace(4)
	PrintRowSpace(3)
	PrintRowSpace(2)

	// Output:
	// *   *
	// *  *
	// * *
	// **
}

func ExamplePrintSquare() {
	PrintSquare(3)
	fmt.Println()
	PrintSquare(5)
	fmt.Println()
	PrintSquare(2)

	// Output:
	// ***
	// ***
	// ***
	//
	// *****
	// *****
	// *****
	// *****
	// *****
	//
	// **
	// **
}

func ExamplePrintEmptySquare() {
	PrintEmptySquare(3)
	fmt.Println()
	PrintEmptySquare(5)
	fmt.Println()
	PrintEmptySquare(2)

	// Output:
	// ***
	// * *
	// ***
	//
	// *****
	// *   *
	// *   *
	// *   *
	// *****
	//
	// **
	// **
}

func ExamplePrintCustomSquare() {
	PrintCustomSquare(3, "A", "B")
	fmt.Println()
	PrintCustomSquare(5, "A", "B")
	fmt.Println()
	PrintCustomSquare(2, "A", "B")

	// Output:
	// AAA
	// ABA
	// AAA
	//
	// AAAAA
	// ABBBA
	// ABBBA
	// ABBBA
	// AAAAA
	//
	// AA
	// AA
}

func ExamplePrintRectangle() {
	PrintRectangle(3, 4)
	fmt.Println()
	PrintRectangle(5, 2)
	fmt.Println()
	PrintRectangle(2, 2)

	// Output:
	// ***
	// ***
	// ***
	// ***
	//
	// *****
	// *****
	//
	// **
	// **
}

func ExamplePrintTriangle() {
	PrintTriangle(3)
	fmt.Println()
	PrintTriangle(5)
	fmt.Println()
	PrintTriangle(2)

	// Output:
	// ***
	// **
	// *
	//
	// *****
	// ****
	// ***
	// **
	// *
	//
	// **
	// *
}
