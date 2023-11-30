package structs

import "fmt"

func ExampleEntry() {
	e1 := Entry{
		de: "Holz",
		en: "wood",
	}

	e2 := Entry{
		de: "Wald",
		en: "wood",
	}

	e3 := Entry{
		de: "Wald",
		en: "forest",
	}

	fmt.Println(e1.Translate())
	fmt.Println(e2.Translate())
	fmt.Println(e3.Translate())

	// Output:
	// wood
	// wood
	// forest
}

func ExampleDictionary() {
	d1 := Dictionary{}

	fmt.Println(d1)

	d1.Add("Holz", "wood")
	d1.Add("Wald", "wood")
	d1.Add("Wald", "forest")

	fmt.Println(d1)

	// Output:
}
