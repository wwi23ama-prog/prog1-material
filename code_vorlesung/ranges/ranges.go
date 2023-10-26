package ranges

import "fmt"

func Loop_normal() {
	l1 := []string{
		"A", "B", "C", "D",
	}

	// "normale" for-Schleife
	for i := 0; i < len(l1); i++ {
		fmt.Println(l1[i])
	}
}

func RangeLoop() {
	l1 := []string{
		"A", "B", "C", "D",
	}

	for i, el := range l1 {
		fmt.Printf("%v: %v\n", i, el)
		//fmt.Println(i, ":", el)
	}

	for i := range l1 {
		fmt.Printf("%v\n", i)
	}

	for _, el := range l1 {
		fmt.Printf("%v\n", el)
	}
}

type Row []string

func RangeLoopRow() {
	r1 := Row{
		"A", "B", "C", "D",
	}

	for pos, v := range r1 {
		fmt.Println(pos, v)
	}
}

func CreateRow() {

	// r1 := Row{}
	r1 := make(Row, 4)
	//r1 := make([]string, 0)

	for i := 0; i < 4; i++ {
		r1 = append(r1, fmt.Sprintf("%v", i))
	}

	fmt.Println(r1)
}

func CreateIntSlice() {
	s1 := make([]int, 4)

	fmt.Println(s1)
	s1[3] = 42
	fmt.Println(s1)
	s1 = append(s1, 38)
	fmt.Println(s1)
}
