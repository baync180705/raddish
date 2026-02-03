package store

import (
	"fmt"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) GET(dbName string, k string) (string, int) {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	r.mu.Unlock()

	if !ok {
		fmt.Println(msg.ErrorDBNotFound)
		return "<INVALID>", 0
	}

	registry.mu.RLock()
	defer registry.mu.RUnlock()

	val, ok := registry.register[k]
	if !ok {
		fmt.Println(msg.ErrorKeyNotFound)
		return "<INVALID>", 0
	}
	return val, 1
}