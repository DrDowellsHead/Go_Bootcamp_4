package domain

const (
	Empty    = 0
	Human    = 1
	Computer = 2
)

// Board — поле 3×3. При присваивании все девять клеток копируются.
type Board [3][3]int

type Game struct {
	ID    string
	Board Board
}
