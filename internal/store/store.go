package store

import (
	"fmt"
	"sync"

	"github.com/baync180705/raddish/internal/msg"
)

type Raddish struct {
	mu sync.Mutex
	db map[string]*registry
}

type registry struct {
	mu       sync.RWMutex
	register map[string]string
}

func INIT() *Raddish {
	fmt.Println(msg.InfoRaddishInitialized)
	return &Raddish{
		db: make(map[string]*registry),
	}
}
