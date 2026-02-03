package store

import (
	"strings"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) LISTDB() (string, int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	keys := make([]string, 0, len(r.db))

	for k := range r.db {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return msg.ErrorNoDBsAvailable, 0
	}

	res := strings.Join(keys, "\n")
	return res, 1
}