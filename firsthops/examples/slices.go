// Basic examples and patterns with slices
package examples 

import (
	"fmt"
)

var (
	buffer [256]byte
)

func addOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}

// It does just what its name implies, iterating over the indices of a slice (using a for range loop), incrementing its elements.
func IncrementSliceElements() {
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }
    fmt.Println("before", slice)
    addOneToEachElement(slice)
    fmt.Println("after", slice)
}

