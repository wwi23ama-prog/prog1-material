package structs

import "fmt"

type Position struct {
	x int
	y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}
