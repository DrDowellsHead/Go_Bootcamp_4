package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrInvalidMove  = errors.New("invalid move")
	ErrGameFinished = errors.New("game already finished")
	ErrNotFound     = errors.New("game not found")
)

type GameService struct {
	repository Repository
	// The lock protects the full read/validate/move/save transaction.
	// sync.Map alone protects individual operations, not their sequence.
	mu sync.Mutex
}

func NewService(repository Repository) Service {
	return &GameService{repository: repository}
}

func (s *GameService) Create() (Game, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return Game{}, fmt.Errorf("generate UUID: %w", err)
	}
	raw[6] = (raw[6] & 0x0f) | 0x40
	raw[8] = (raw[8] & 0x3f) | 0x80
	game := Game{
		ID: fmt.Sprintf("%x-%x-%x-%x-%x", raw[0:4], raw[4:6], raw[6:8], raw[8:10], raw[10:16]),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.repository.Save(game); err != nil {
		return Game{}, err
	}
	return game, nil
}

func (s *GameService) Get(id string) (Game, error) {
	game, ok, err := s.repository.Get(id)
	if err != nil {
		return Game{}, err
	}
	if !ok {
		return Game{}, ErrNotFound
	}
	return game, nil
}

func (s *GameService) Play(id string, board Board) (Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	previous, found, err := s.repository.Get(id)
	if err != nil {
		return Game{}, err
	}
	if !found {
		// A client may also start by choosing its own UUID and first move.
		previous = Game{ID: id}
	}
	submitted := Game{ID: id, Board: board}
	if err := s.ValidateBoard(previous, submitted); err != nil {
		return Game{}, err
	}

	updated, err := s.NextMove(submitted)
	if err != nil {
		return Game{}, err
	}
	if err := s.repository.Save(updated); err != nil {
		return Game{}, err
	}
	return updated, nil
}

func (s *GameService) ValidateBoard(previous, submitted Game) error {
	if previous.ID != submitted.ID {
		return fmt.Errorf("%w: game IDs do not match", ErrInvalidMove)
	}
	if finished, _ := s.IsFinished(previous); finished {
		return ErrGameFinished
	}

	oldHuman, oldComputer := 0, 0
	newHuman, newComputer := 0, 0
	changes := 0

	for row := range previous.Board {
		for col, old := range previous.Board[row] {
			current := submitted.Board[row][col]
			if old < Empty || old > Computer || current < Empty || current > Computer {
				return fmt.Errorf("%w: cells must contain only 0, 1 or 2", ErrInvalidMove)
			}
			if old == Human {
				oldHuman++
			} else if old == Computer {
				oldComputer++
			}
			if current == Human {
				newHuman++
			} else if current == Computer {
				newComputer++
			}
			if old != current {
				if old != Empty || current != Human {
					return fmt.Errorf("%w: previous moves cannot be changed; only a free cell can become 1", ErrInvalidMove)
				}
				changes++
			}
		}
	}

	if oldHuman != oldComputer || changes != 1 || newHuman != oldHuman+1 || newComputer != oldComputer {
		return fmt.Errorf("%w: submit exactly one new human move", ErrInvalidMove)
	}
	return nil
}

func (s *GameService) NextMove(game Game) (Game, error) {
	if finished, _ := s.IsFinished(game); finished {
		return game, nil
	}
	board, ok := bestMove(game.Board)
	if !ok {
		return game, fmt.Errorf("%w: board has no free cells", ErrInvalidMove)
	}
	game.Board = board
	return game, nil
}

func (s *GameService) IsFinished(game Game) (bool, int) {
	return result(game.Board)
}
