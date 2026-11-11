package domain

// Service describes game behavior without mentioning HTTP or storage.
type Service interface {
	Create() (Game, error)
	Get(id string) (Game, error)
	Play(id string, board Board) (Game, error)
	NextMove(game Game) (Game, error)
	ValidateBoard(previous, submitted Game) error
	IsFinished(game Game) (finished bool, winner int)
}
