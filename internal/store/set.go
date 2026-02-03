package store

import (
	"fmt"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) SET(dbName string, k string, v string) int {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	r.mu.Unlock()

	if !ok {
		fmt.Println(msg.ErrorDBNotFound)
		return 0
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.register[k] = v
	return 1
}