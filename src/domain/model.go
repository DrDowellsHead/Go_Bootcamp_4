package domain

const (
	Empty    = 0
	Human    = 1
	Computer = 2
)

// Board is a value type: copying it copies all nine cells.
type Board [3][3]int

type Game struct {
	ID    string
	Board Board
}
