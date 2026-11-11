package datasource

// GameRecord is the storage representation, independent of the domain model.
type GameRecord struct {
	ID    string
	Cells [3][3]int
}
