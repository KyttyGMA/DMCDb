# Memory Pledge Algorithm

## 1D to 2D Formula

The algorithm uses the following formula to map a 1D index to a 2D index:

**index = i × cols + j**

Where:
- `index` is the 1D index.
- `i` is the row index in the 2D array.
- `j` is the column index in the 2D array.
- `cols` is the number of columns in the 2D array.

## 1D to 0D Algorithm

The following Go code demonstrates how to convert a 1D array to a single scalar value using a SHA-256 hash function.

```go
package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Function to convert a 1D array to a single scalar value using SHA-256 hash
func memoryPledge1Dto0D(data []int) [32]byte {
	// Convert the 1D array to a byte slice
	dataBytes := make([]byte, len(data)*8) // Assuming int is 64-bit (8 bytes)
	for i, num := range data {
		binary.LittleEndian.PutUint64(dataBytes[i*8:(i+1)*8], uint64(num))
	}

	// Compute the SHA-256 hash of the byte slice
	hash := sha256.Sum256(dataBytes)

	// Return the hash value
	return hash
}

func main() {
	// Example 1D array
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Perform memory pledge from 1D to 0D
	hashValue := memoryPledge1Dto0D(data)

	// Output the hash value
	fmt.Printf("Hash Value: %x\n", hashValue)
}
