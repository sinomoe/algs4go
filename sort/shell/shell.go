package shell

import (
	"algs4go/interfaces"
)

// Shell sort
func Shell(datas interfaces.Sortable) interfaces.Sortable {
	N := datas.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j > h-1 && datas.Less(j, j-h); j -= h {
				datas.Swap(j, j-h)
			}
		}
		h /= 3
	}

	return datas
}
