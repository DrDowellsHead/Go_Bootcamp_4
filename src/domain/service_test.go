package domain

import (
	"errors"
	"testing"
)

type testRepository struct {
	items map[string]Game
}

func (r *testRepository) Save(game Game) error {
	r.items[game.ID] = game
	return nil
}
func (r *testRepository) Get(id string) (Game, bool, error) {
	game, ok := r.items[id]
	return game, ok, nil
}

func TestMinimaxTakesWin(t *testing.T) {
	game := Game{ID: "win", Board: Board{
		{Computer, Computer, Empty},
		{Human, Human, Empty},
		{Empty, Empty, Empty},
	}}
	got, err := NewService(&testRepository{items: map[string]Game{}}).NextMove(game)
	if err != nil {
		t.Fatal(err)
	}
	if got.Board[0][2] != Computer {
		t.Fatalf("computer should win immediately, got %v", got.Board)
	}
}

func TestMinimaxBlocksHuman(t *testing.T) {
	game := Game{ID: "block", Board: Board{
		{Human, Human, Empty},
		{Empty, Computer, Empty},
		{Empty, Empty, Empty},
	}}
	got, err := NewService(&testRepository{items: map[string]Game{}}).NextMove(game)
	if err != nil {
		t.Fatal(err)
	}
	if got.Board[0][2] != Computer {
		t.Fatalf("computer should block an immediate loss, got %v", got.Board)
	}
}

func TestPlayPreservesPreviousMoves(t *testing.T) {
	repo := &testRepository{items: make(map[string]Game)}
	service := NewService(repo)
	first := Board{{Human, Empty, Empty}}
	game, err := service.Play("first", first)
	if err != nil {
		t.Fatal(err)
	}
	if game.Board[0][0] != Human {
		t.Fatal("human move was lost")
	}
	if _, err = service.Play("first", Board{{Empty, Human, Empty}}); !errors.Is(err, ErrInvalidMove) {
		t.Fatalf("changed previous move must fail, got %v", err)
	}
	if _, err = service.Play("first", game.Board); !errors.Is(err, ErrInvalidMove) {
		t.Fatalf("missing new move must fail, got %v", err)
	}
}

func TestInvalidCellsAndMultipleMoves(t *testing.T) {
	service := NewService(&testRepository{items: make(map[string]Game)})
	for _, board := range []Board{
		{{3, 0, 0}},
		{{Human, Human, Empty}},
		{{Computer, Empty, Empty}},
	} {
		if _, err := service.Play("bad", board); !errors.Is(err, ErrInvalidMove) {
			t.Fatalf("board %v unexpectedly accepted: %v", board, err)
		}
	}
}

func TestCreateAndIndependentGames(t *testing.T) {
	service := NewService(&testRepository{items: make(map[string]Game)})
	a, err := service.Create()
	if err != nil {
		t.Fatal(err)
	}
	b, err := service.Create()
	if err != nil {
		t.Fatal(err)
	}
	if a.ID == b.ID || len(a.ID) != 36 {
		t.Fatalf("invalid generated UUIDs: %q and %q", a.ID, b.ID)
	}
	if _, err := service.Play(a.ID, Board{{Human, Empty, Empty}}); err != nil {
		t.Fatal(err)
	}
	storedB, err := service.Get(b.ID)
	if err != nil || storedB.Board != (Board{}) {
		t.Fatalf("second game changed: %+v, %v", storedB, err)
	}
	if _, err := service.Get("unknown"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestFinishedGame(t *testing.T) {
	repo := &testRepository{items: map[string]Game{
		"done": {ID: "done", Board: Board{
			{Human, Human, Human},
			{Computer, Computer, Empty},
			{Empty, Empty, Empty},
		}},
	}}
	service := NewService(repo)
	finished, winner := service.IsFinished(repo.items["done"])
	if !finished || winner != Human {
		t.Fatalf("expected human victory, got %v %v", finished, winner)
	}
	if _, err := service.Play("done", repo.items["done"].Board); !errors.Is(err, ErrGameFinished) {
		t.Fatalf("finished game should be immutable: %v", err)
	}
	draw := Game{Board: Board{
		{Human, Computer, Human},
		{Human, Computer, Computer},
		{Computer, Human, Human},
	}}
	if finished, winner := service.IsFinished(draw); !finished || winner != Empty {
		t.Fatalf("expected draw, got finished=%v winner=%d", finished, winner)
	}
}
