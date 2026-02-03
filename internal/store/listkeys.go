package store

import (
	"strings"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) LISTKEYS(dbName string) (string, int) {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	if !ok {
		r.mu.Unlock()
		return msg.ErrorDBNotFound, 0
	}
	r.mu.Unlock()

	registry.mu.RLock()
	defer registry.mu.RUnlock()

	keys := make([]string, 0, len(registry.register))

	for k := range registry.register {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return msg.ErrorNoKeysInDB, 0
	}

	res := strings.Join(keys, "\n")
	return res, 1
}