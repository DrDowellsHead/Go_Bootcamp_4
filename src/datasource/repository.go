package datasource

import "github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"

type Repository struct {
	storage *Storage
}

func NewRepository(storage *Storage) *Repository {
	return &Repository{storage: storage}
}

func (r *Repository) Save(game domain.Game) error {
	r.storage.games.Store(game.ID, toRecord(game))
	return nil
}

func (r *Repository) Get(id string) (domain.Game, bool, error) {
	value, found := r.storage.games.Load(id)
	if !found {
		return domain.Game{}, false, nil
	}
	return toDomain(value.(GameRecord)), true, nil
}

var _ domain.Repository = (*Repository)(nil)
