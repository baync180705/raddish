package store

import (
	"fmt"

	"github.com/baync180705/raddish/internal/msg"
)

func (r *Raddish) CREATE(k string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.db[k]; ok {
		fmt.Println(msg.ErrorDBAlreadyExists)
		return 0
	}
	r.db[k] = &registry{
		register: make(map[string]string),
	}
	return 1
}
