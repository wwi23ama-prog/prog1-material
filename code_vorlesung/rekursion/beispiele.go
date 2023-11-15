package rekursion

import "fmt"

func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)
}
