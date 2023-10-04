package main

import "fmt"

func main() {
	var n int
	fmt.Print("Bitte eine Zahl eingeben: ")
	fmt.Scan(&n)

	if n != 42 {
		fmt.Println("Das war falsch!")
		return
	}
	fmt.Print("Ihre Lieblingszahl: ", n)
}
