package main

import (
	"fmt"
	"sync"
)

// Metadata structure for the projection of the full data
type Metadata struct {
	PrimaryKey   string
	SecondaryKey string
	// Add other fields you want to project
}

// DualMemoryDB structure for the database
type DualMemoryDB struct {
	Primary   map[string]Metadata // Primary hashmap
	Secondary map[string]string   // Secondary hashmap
	mu        sync.Mutex          // Mutex for protecting concurrent access
}

// Constructor for the database
func NewDualMemoryDB() *DualMemoryDB {
	return &DualMemoryDB{
		Primary:   make(map[string]Metadata),
		Secondary: make(map[string]string),
	}
}

// Function to store data
func (db *DualMemoryDB) StoreData(primaryKey, secondaryKey, data string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Create the metadata structure
	metadata := Metadata{
		PrimaryKey:   primaryKey,
		SecondaryKey: secondaryKey,
		// Add other fields you want to project
	}

	// Store in primary hashmap
	db.Primary[primaryKey] = metadata

	// Store in secondary hashmap
	db.Secondary[secondaryKey] = data
}

// Function to retrieve metadata by primary key
func (db *DualMemoryDB) RetrieveByPrimaryKey(primaryKey string) (Metadata, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	metadata, exists := db.Primary[primaryKey]
	return metadata, exists
}

// Function to retrieve full data by secondary key
func (db *DualMemoryDB) RetrieveBySecondaryKey(secondaryKey string) (string, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, exists := db.Secondary[secondaryKey]
	return data, exists
}

func main() {
	// Create a new instance of the database
	db := NewDualMemoryDB()

	// Example data
	db.StoreData("pk1", "sk1", "data1")
	db.StoreData("pk2", "sk2", "data2")

	// Retrieve metadata by primary key
	metadata, exists := db.RetrieveByPrimaryKey("pk1")
	if exists {
		fmt.Println("Metadata by Primary Key:", metadata)
	} else {
		fmt.Println("Primary Key not found")
	}

	// Retrieve full data by secondary key
	data, exists := db.RetrieveBySecondaryKey("sk2")
	if exists {
		fmt.Println("Full Data by Secondary Key:", data)
	} else {
		fmt.Println("Secondary Key not found")
	}
}
