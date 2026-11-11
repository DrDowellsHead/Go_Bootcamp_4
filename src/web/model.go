package web

// Game is the public JSON contract. Human is 1, computer is 2, empty is 0.
type Game struct {
	ID     string    `json:"id"`
	Board  [3][3]int `json:"board"`
	Status string    `json:"status,omitempty"`
	Winner int       `json:"winner"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
