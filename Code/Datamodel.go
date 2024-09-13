package main

import "time"

// Estructura para los metadatos
type Metadata struct {
	ID         string
	Size       int
	CreatedAt  time.Time
	ModifiedAt time.Time
	Hash       string
	Tags       []string
}

// Estructura de la base de datos
type DualMemoryDB struct {
	Primary   map[string]Metadata // Capa primaria (RAM)
	Secondary map[string]string    // Capa secundaria (almacenamiento)
}

// Constructor para la base de datos
func NewDualMemoryDB() *DualMemoryDB {
	return &DualMemoryDB{
		Primary:   make(map[string]Metadata),
		Secondary: make(map[string]string),
	}
}
