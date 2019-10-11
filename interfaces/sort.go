package interfaces

import "sort"

// Sortable defines the interface of comparable items
type Sortable interface {
	sort.Interface
	// Move(dest, src int) bool
	// Get(i int) (interface{}, error)
	// Set(i int, v interface{}) (interface{}, error)
}
