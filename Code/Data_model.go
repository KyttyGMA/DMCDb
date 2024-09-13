package main

import (
	"sync"
	"time"
)

// Estructura de la base de datos
type DualMemoryDB struct {
	Primary   map[string]Metadata
	Secondary map[string]string
	mu        sync.Mutex // Mutex para proteger el acceso concurrente
}

// Constructor para la base de datos
func NewDualMemoryDB() *DualMemoryDB {
	return &DualMemoryDB{
		Primary:   make(map[string]Metadata),
		Secondary: make(map[string]string),
	}
}
