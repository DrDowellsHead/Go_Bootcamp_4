package web

// Game — формат ответа API: 0 — пустая клетка, 1 — игрок, 2 — компьютер.
type Game struct {
	ID     string    `json:"id"`
	Board  [3][3]int `json:"board"`
	Status string    `json:"status,omitempty"`
	Winner int       `json:"winner"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
