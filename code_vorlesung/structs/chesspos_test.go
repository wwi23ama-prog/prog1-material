package structs

import "fmt"

func ExampleChessPos() {
	p1 := ChessPos{row: 3, col: "B"}

	fmt.Println(p1)

	fmt.Println(p1.String())

	// Output:
}

func ExampleChessPos_Add() {
	p1 := ChessPos{row: 3, col: "B"}
	p2 := ChessPos{row: 5, col: "C"}

	p3 := p1.Add(p2)

	fmt.Println(p3)

	// Output:
}
