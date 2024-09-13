package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Función para calcular el hash de un dato
func hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Función para insertar datos
func (db *DualMemoryDB) Insert(key string, value string) {
	hashedKey := hash(key)

	// Crear metadatos
	metadata := Metadata{
		ID:         hashedKey,
		Size:       len(value),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Hash:       hash(value),
		Tags:       []string{"general"},
	}

	// Insertar en la capa primaria
	db.Primary[hashedKey] = metadata

	// Insertar el dato completo en la capa secundaria
	db.Secondary[hashedKey] = value
}

// Función para recuperar datos completos
func (db *DualMemoryDB) GetFullData(key string) string {
	hashedKey := hash(key)
	return db.Secondary[hashedKey]
}

// Función para recuperar metadatos
func (db *DualMemoryDB) GetMetadata(key string) Metadata {
	hashedKey := hash(key)
	return db.Primary[hashedKey]
}
