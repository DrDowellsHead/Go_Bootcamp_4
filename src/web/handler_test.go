package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"
)

type memoryRepository struct {
	games map[string]domain.Game
}

func (m *memoryRepository) Save(game domain.Game) error {
	m.games[game.ID] = game
	return nil
}

func (m *memoryRepository) Get(id string) (domain.Game, bool, error) {
	game, ok := m.games[id]
	return game, ok, nil
}

func TestHTTPGameLifecycle(t *testing.T) {
	service := domain.NewService(&memoryRepository{games: make(map[string]domain.Game)})
	handler := NewHandler(service).Routes()
	create := httptest.NewRecorder()
	handler.ServeHTTP(create, httptest.NewRequest(http.MethodPost, "/game", nil))
	if create.Code != http.StatusCreated {
		t.Fatalf("create status = %d: %s", create.Code, create.Body)
	}
	var newGame Game
	if err := json.Unmarshal(create.Body.Bytes(), &newGame); err != nil {
		t.Fatal(err)
	}
	newGame.Board[0][0] = domain.Human
	payload, _ := json.Marshal(newGame)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/game/"+newGame.ID, bytes.NewReader(payload)))
	if response.Code != http.StatusOK {
		t.Fatalf("play status = %d: %s", response.Code, response.Body)
	}
	var after Game
	if err := json.Unmarshal(response.Body.Bytes(), &after); err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, row := range after.Board {
		for _, cell := range row {
			if cell == domain.Computer {
				count++
			}
		}
	}
	if count != 1 {
		t.Fatalf("expected exactly one computer move, got %v", after.Board)
	}

	replayed := httptest.NewRecorder()
	handler.ServeHTTP(replayed, httptest.NewRequest(http.MethodPost, "/game/"+newGame.ID, bytes.NewReader(payload)))
	if replayed.Code != http.StatusBadRequest {
		t.Fatalf("replay should be rejected: %d", replayed.Code)
	}
}

func TestHTTPInvalidJSON(t *testing.T) {
	service := domain.NewService(&memoryRepository{games: make(map[string]domain.Game)})
	handler := NewHandler(service).Routes()
	id := "c2b49acc-5f27-4f70-9b9f-f9e836fdabef"
	for _, body := range []string{`{"board":[[1,0,0]]}`, `not json`, `{"board":[[1,0,0],[0,0,0],[0,0,0]]}{}`} {
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/game/"+id, bytes.NewBufferString(body)))
		if response.Code != http.StatusBadRequest {
			t.Fatalf("body %q: status %d", body, response.Code)
		}
	}
}
