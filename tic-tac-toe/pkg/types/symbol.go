package types

import "fmt"

type Symbol int

const (
	Empty Symbol = iota
	X
	O
)

func (s Symbol) String() string {
	switch s {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "."
	}
}

type Position struct {
	Row, Col int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Row, p.Col)
}
