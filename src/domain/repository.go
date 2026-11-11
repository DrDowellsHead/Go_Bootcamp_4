package domain

// Repository задаёт методы для сохранения и получения игры.
type Repository interface {
	Save(Game) error
	Get(id string) (Game, bool, error)
}
