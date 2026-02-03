package store

import (
	"fmt"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) DEL(dbName string, k string) int {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	r.mu.Unlock()

	if !ok {
		fmt.Println(msg.ErrorDBNotFound)
		return 0
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()

	_, ok = registry.register[k]
	if !ok {
		fmt.Println(msg.ErrorKeyNotFoundDel)
		return 0
	}

	delete(registry.register, k)
	return 1
}