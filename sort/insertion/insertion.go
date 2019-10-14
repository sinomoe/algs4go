package insertion

import (
	"algs4go/interfaces"
)

// Insertion implements basic insertion sort algorithm
func Insertion(datas interfaces.Sortable) interfaces.Sortable {
	N := datas.Len()
	for i := 1; i < N; i++ {
		for j := i; j > 0 && datas.Less(j, j-1); j-- {
			datas.Swap(j, j-1)
		}
	}
	return datas
}
