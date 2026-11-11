package datasource

import "sync"

// Storage owns a thread-safe map shared by every repository.
type Storage struct {
	games sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}
