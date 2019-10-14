package quick

import "algs4go/interfaces"

// Quick sort
func Quick(datas interfaces.Sortable) interfaces.Sortable {
	quickSort(datas, 0, datas.Len()-1)
	return datas
}

func quickSort(datas interfaces.Sortable, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(datas, lo, hi)
	quickSort(datas, lo, j-1)
	quickSort(datas, j+1, hi)
}

func partition(datas interfaces.Sortable, lo, hi int) int {
	i, j := lo, hi+1
	loIdx := lo
	for {
		for i = i + 1; datas.Less(i, loIdx); i++ {
			if i == hi {
				break
			}
		}
		for j = j - 1; datas.Less(loIdx, j); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		if i == lo {
			loIdx = j
		}
		if j == lo {
			loIdx = i
		}
		datas.Swap(i, j)
	}
	datas.Swap(lo, j)

	return j
}
