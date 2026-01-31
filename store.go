package main

import (
	"fmt"
	"sync"
)

type Raddish struct {
	mu      sync.Mutex
	db      map[string]*registry
}

type registry struct {
	mu       sync.RWMutex
	register map[string]string
}

func INIT() *Raddish {
	fmt.Println("Raddish initialized successfully")
	return &Raddish{
		db: make(map[string]*registry),
	}
}

func (r *Raddish) CREATE(k string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.db[k]; ok {
		fmt.Println("Cannot create an existing key")
		return 0
	}
	r.db[k] = &registry{
		register: make(map[string]string),
	}
	return 1
}

func (r *Raddish) SET(dbName string, k string, v string) int {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	r.mu.Unlock()

	if !ok {
		fmt.Println("Given DB does not exist")
		return 0
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.register[k] = v

	return 1
}

func (r *Raddish) GET(dbName string, k string) (string, int) {
	r.mu.Lock()
	registry, ok := r.db[dbName] 
	r.mu.Unlock()

	if !ok {
		fmt.Println("Given DB does not exist")
		return "<INVALID>", 0
	}

	registry.mu.RLock()
	defer registry.mu.RUnlock()
	
	val, ok := registry.register[k]; if !ok {
		fmt.Println("Given key is unavailable")
		return "<INVALID>", 0
	}
	return val, 1
}

func (r *Raddish) DEL(dbName string, k string) int {
	r.mu.Lock()
	registry, ok := r.db[dbName]
	r.mu.Unlock()

	if !ok {
		fmt.Println("Given DB does not exist")
		return 0
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()

	_, ok = registry.register[k]; if !ok {
		fmt.Println("Given key not fonund")
		return 0
	}

	delete(registry.register, k)
	return 1
}
