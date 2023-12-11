package fragestunde

import "fmt"

func ExampleSprint() {
	fmt.Println(42)

	s1 := fmt.Sprint(42)
	fmt.Println(s1)

	fmt.Printf("%d\n", 42)

	s2 := fmt.Sprintf("%d", 42)
	fmt.Println(s2)

	s3 := fmt.Sprintf("%d %f %v %T %t %s",
		42,   // %d: Als Dezimalzahl
		42.0, // %f: Als Fließkommazahl
		42,   // %v: Als Wert
		42,   // %T: Als Typ
		true, // %t: Als Wahrheitswert
		"42", // %s: Als String
	)

	fmt.Println(s3)

	// Output:
}

// Foo2 berechnet den größten gemeinsamen Teiler von x und y.
func Foo2(x, y int) int {
	if x == y {
		return x
	}
	if x > y {
		return Foo2(x-y, y)
	}
	return Foo2(x, y-x)
}
