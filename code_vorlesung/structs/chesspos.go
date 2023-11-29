package structs

import "fmt"

type ChessPos struct {
	row int
	col string
}

// Liefert die Position im Format B3.
func (p ChessPos) String() string {
	return fmt.Sprintf("%s%d", p.col, p.row)
}

func (p ChessPos) Add(q ChessPos) ChessPos {
	return ChessPos{
		row: p.row + q.row,
		col: p.col + q.col,
	}
}
