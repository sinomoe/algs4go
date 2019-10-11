package selection

import "algs4go/interfaces"

// Selection sort
func Selection(datas interfaces.Sortable) interfaces.Sortable {
	N := datas.Len()
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if datas.Less(j, min) {
				min = j
			}
		}
		datas.Swap(i, min)
	}
	return datas
}
