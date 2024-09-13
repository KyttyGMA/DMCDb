package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/snappy"
)

// Metadata structure for the projection of the full data
type Metadata struct {
	ID         string
	Size       int
	CreatedAt  string
	ModifiedAt string
	Hash       string
	Tags       []string
}

// DualMemoryDB structure for the database
type DualMemoryDB struct {
	Primary   map[string]Metadata
	Secondary map[string]string
}

// Constructor for the database
func NewDualMemoryDB() *DualMemoryDB {
	return &DualMemoryDB{
		Primary:   make(map[string]Metadata),
		Secondary: make(map[string]string),
	}
}

// Hash function to generate a hash from a string
func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

// Compress function to compress data using Snappy
func compress(data []byte) []byte {
	var buf bytes.Buffer
	writer := snappy.NewBufferedWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Decompress function to decompress data using Snappy
func decompress(data []byte) []byte {
	return snappy.Decode(nil, data)
}

// SaveCompressedData saves compressed data to a local file
func SaveCompressedData(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadCompressedData loads compressed data from a local file
func LoadCompressedData(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Insert function to insert data into the database
func (db *DualMemoryDB) Insert(key string, value string) {
	hashedKey := hash(key)

	// Create metadata
	metadata := Metadata{
		ID:         hashedKey,
		Size:       len(value),
		CreatedAt:  time.Now().Format(time.RFC3339),
		ModifiedAt: time.Now().Format(time.RFC3339),
		Hash:       hash(value),
		Tags:       []string{"general"},
	}

	// Serialize metadata to JSON
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		panic(err)
	}

	// Compress metadata
	compressedMetadata := compress(metadataBytes)

	// Insert compressed metadata into primary layer
	db.Primary[hashedKey] = string(compressedMetadata)

	// Compress the full data
	compressedValue := compress([]byte(value))

	// Save the compressed data to a local file
	filename := filepath.Join("data", hashedKey+".snappy")
	err = SaveCompressedData(filename, compressedValue)
	if err != nil {
		panic(err)
	}

	// Store the filename in the secondary layer
	db.Secondary[hashedKey] = filename
}

// Retrieve function to retrieve data from the database
func (db *DualMemoryDB) Retrieve(key string) (string, bool) {
	hashedKey := hash(key)

	// Retrieve metadata from primary layer
	compressedMetadata, exists := db.Primary[hashedKey]
	if !exists {
		return "", false
	}

	// Decompress metadata
	metadataBytes := decompress([]byte(compressedMetadata))
	var metadata Metadata
	err := json.Unmarshal(metadataBytes, &metadata)
	if err != nil {
		panic(err)
	}

	// Retrieve the filename from the secondary layer
	filename, exists := db.Secondary[hashedKey]
	if !exists {
		return "", false
	}

	// Load the compressed data from the local file
	compressedData, err := LoadCompressedData(filename)
	if err != nil {
		panic(err)
	}

	// Decompress the data
	data := decompress(compressedData)

	return string(data), true
}

func main() {
	// Create a new instance of the database
	db := NewDualMemoryDB()

	// Ensure the data directory exists
	err := os.MkdirAll("data", 0755)
	if err != nil {
		panic(err)
	}

	// Example data
	db.Insert("key1", "value1")
	db.Insert("key2", "value2")

	// Retrieve data
	value, exists := db.Retrieve("key1")
	if exists {
		fmt.Println("Retrieved Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
