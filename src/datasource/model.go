package datasource

// GameRecord — модель игры для хранения данных.
type GameRecord struct {
	ID    string
	Cells [3][3]int
}
