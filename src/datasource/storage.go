package datasource

import "sync"

// Storage хранит игры в потокобезопасной sync.Map.
type Storage struct {
	games sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}
