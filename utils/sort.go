package utils

import (
	"algs4go/interfaces"
	"sort"
)

// IsSorted returns whether input datas is sorted
func IsSorted(in interfaces.Sortable) bool {
	return sort.IsSorted(in)
}
