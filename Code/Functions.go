package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

// Función para calcular el hash de un dato
func hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Función para insertar datos
func (db *DualMemoryDB) Insert(key string, value string) {
	if len(key) == 0 || len(value) == 0 {
		fmt.Println("Error: La clave y el valor no pueden estar vacíos.")
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()

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

// Función para actualizar datos
func (db *DualMemoryDB) Update(key string, newValue string) {
	if len(key) == 0 || len(newValue) == 0 {
		fmt.Println("Error: La clave y el valor no pueden estar vacíos.")
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	hashedKey := hash(key)

	// Verificar si existe el dato
	if _, exists := db.Primary[hashedKey]; !exists {
		fmt.Println("Error: La clave no existe.")
		return
	}

	// Actualizar metadatos
	metadata := db.Primary[hashedKey]
	metadata.Size = len(newValue)
	metadata.ModifiedAt = time.Now()
	metadata.Hash = hash(newValue) // Actualizar hash

	// Guardar cambios
	db.Primary[hashedKey] = metadata
	db.Secondary[hashedKey] = newValue // Actualizar en la capa secundaria
}

// Función para eliminar datos
func (db *DualMemoryDB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	hashedKey := hash(key)

	// Eliminar de ambas capas
	delete(db.Primary, hashedKey)
	delete(db.Secondary, hashedKey)
}

// Función para recuperar datos completos
func (db *DualMemoryDB) GetFullData(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	hashedKey := hash(key)
	return db.Secondary[hashedKey]
}

// Función para recuperar metadatos
func (db *DualMemoryDB) GetMetadata(key string) Metadata {
	db.mu.Lock()
	defer db.mu.Unlock()

	hashedKey := hash(key)
	return db.Primary[hashedKey]
}

// Función para buscar datos por etiqueta
func (db *DualMemoryDB) SearchByTag(tag string) []Metadata {
	db.mu.Lock()
	defer db.mu.Unlock()

	var results []Metadata
	for _, metadata := range db.Primary {
		for _, t := range metadata.Tags {
			if t == tag {
				results = append(results, metadata)
			}
		}
	}
	return results
}

// Función para guardar la capa secundaria en un archivo
func (db *DualMemoryDB) SaveToFile(filename string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := json.Marshal(db.Secondary)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// Función para cargar la capa secundaria desde un archivo
func (db *DualMemoryDB) LoadFromFile(filename string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &db.Secondary)
}
