package domain

// Repository hides the storage implementation from the game rules.
type Repository interface {
	Save(Game) error
	Get(id string) (Game, bool, error)
}
