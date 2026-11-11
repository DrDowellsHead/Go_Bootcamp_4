package datasource

import (
	"sync"
	"testing"

	"github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"
)

func TestRepositoryRoundTrip(t *testing.T) {
	repo := NewRepository(NewStorage())
	want := domain.Game{ID: "one", Board: domain.Board{{1, 2, 0}}}
	if err := repo.Save(want); err != nil {
		t.Fatal(err)
	}
	got, ok, err := repo.Get(want.ID)
	if err != nil || !ok || got != want {
		t.Fatalf("round trip: got %+v, exists %t, error %v", got, ok, err)
	}
	_, ok, err = repo.Get("missing")
	if err != nil || ok {
		t.Fatalf("missing record: found=%t error=%v", ok, err)
	}
}

func TestConcurrentStorage(t *testing.T) {
	repo := NewRepository(NewStorage())
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			game := domain.Game{ID: string(rune(n + 100))}
			if err := repo.Save(game); err != nil {
				t.Error(err)
			}
			if _, ok, err := repo.Get(game.ID); err != nil || !ok {
				t.Errorf("concurrent Get: found=%t err=%v", ok, err)
			}
		}(i)
	}
	wg.Wait()
}
