package main

import (
	"testing"
	"time"
)

// TestInsert verifica la inserción de datos en la base de datos.
func TestInsert(t *testing.T) {
	db := NewDualMemoryDB()
	key := "testKey"
	value := "testValue"

	db.Insert(key, value)

	// Verificar la inserción en la capa secundaria
	storedValue := db.GetFullData(key)
	if storedValue != value {
		t.Errorf("Expected value %s, but got %s", value, storedValue)
	}

	// Verificar la inserción en la capa primaria
	metadata := db.GetMetadata(key)
	if metadata.ID == "" {
		t.Error("Expected metadata ID to be non-empty")
	}
	if metadata.Size != len(value) {
		t.Errorf("Expected size %d, but got %d", len(value), metadata.Size)
	}
}

// TestUpdate verifica la actualización de datos en la base de datos.
func TestUpdate(t *testing.T) {
	db := NewDualMemoryDB()
	key := "testKey"
	value := "testValue"
	newValue := "newTestValue"

	db.Insert(key, value)
	db.Update(key, newValue)

	// Verificar la actualización en la capa secundaria
	storedValue := db.GetFullData(key)
	if storedValue != newValue {
		t.Errorf("Expected value %s, but got %s", newValue, storedValue)
	}

	// Verificar la actualización en la capa primaria
	metadata := db.GetMetadata(key)
	if metadata.Size != len(newValue) {
		t.Errorf("Expected size %d, but got %d", len(newValue), metadata.Size)
	}
	if metadata.Hash != hash(newValue) {
		t.Errorf("Expected hash %s, but got %s", hash(newValue), metadata.Hash)
	}
}

// TestDelete verifica la eliminación de datos en la base de datos.
func TestDelete(t *testing.T) {
	db := NewDualMemoryDB()
	key := "testKey"
	value := "testValue"

	db.Insert(key, value)
	db.Delete(key)

	// Verificar que el valor ha sido eliminado de la capa secundaria
	storedValue := db.GetFullData(key)
	if storedValue != "" {
		t.Errorf("Expected value to be empty, but got %s", storedValue)
	}

	// Verificar que los metadatos han sido eliminados de la capa primaria
	metadata := db.GetMetadata(key)
	if metadata.ID != "" {
		t.Error("Expected metadata ID to be empty")
	}
}

// TestSearchByTag verifica la búsqueda por etiqueta en la base de datos.
func TestSearchByTag(t *testing.T) {
	db := NewDualMemoryDB()
	tag := "testTag"
	key1 := "testKey1"
	value1 := "testValue1"
	key2 := "testKey2"
	value2 := "testValue2"

	db.Insert(key1, value1)
	db.Insert(key2, value2)

	// Agregar etiquetas
	metadata1 := db.GetMetadata(key1)
	metadata1.Tags = []string{tag}
	db.Primary[key1] = metadata1

	metadata2 := db.GetMetadata(key2)
	metadata2.Tags = []string{"otherTag"}
	db.Primary[key2] = metadata2

	results := db.SearchByTag(tag)
	if len(results) != 1 || results[0].ID != hash(key1) {
		t.Errorf("Expected to find one entry with tag %s, but found %d", tag, len(results))
	}
}

// TestSaveToFile verifica la funcionalidad de guardar la capa secundaria en un archivo.
func TestSaveToFile(t *testing.T) {
	db := NewDualMemoryDB()
	key := "testKey"
	value := "testValue"

	db.Insert(key, value)
	err := db.SaveToFile("testdata.json")
	if err != nil {
		t.Fatalf("Failed to save to file: %v", err)
	}
}

// TestLoadFromFile verifica la funcionalidad de cargar la capa secundaria desde un archivo.
func TestLoadFromFile(t *testing.T) {
	db := NewDualMemoryDB()
	key := "testKey"
	value := "testValue"

	db.Insert(key, value)
	db.SaveToFile("testdata.json")

	db2 := NewDualMemoryDB()
	err := db2.LoadFromFile("testdata.json")
	if err != nil {
		t.Fatalf("Failed to load from file: %v", err)
	}

	storedValue := db2.GetFullData(key)
	if storedValue != value {
		t.Errorf("Expected value %s, but got %s", value, storedValue)
	}
}
