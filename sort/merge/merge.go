package merge

import "algs4go/interfaces"

// Merge sort. top down merge sort,
// DO NOT USE THIS CODE,
// Because of using sort.Interface to implement this algorithm
// We can only use swap to order the sequence instead of operating on datas direcly
// We have to introduce some loops to calculate how to swap datas
// So it's very slow,
func Merge(datas interfaces.Sortable) interfaces.Sortable {
	mergeSort(datas, 0, datas.Len()-1)
	return datas
}

func mergeSort(datas interfaces.Sortable, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(datas, lo, mid)
	mergeSort(datas, mid+1, hi)
	merge(datas, lo, mid, hi)
}

// merge two ordered subsequence
func merge(datas interfaces.Sortable, lo, mid, hi int) {
	i, j := lo, mid+1
	// In order to use sort.Interface
	// We can only use swap to reorder datas instead of copying data directly
	// so we recorded datas' index and calculate the swap pairs, then swap them
	targetIdx := make([]int, hi-lo+1)
	for k := lo; k <= hi; k++ {
		if i > mid {
			targetIdx[k-lo] = j - lo
			j++
		} else if j > hi {
			targetIdx[k-lo] = i - lo
			i++
		} else if datas.Less(i, j) {
			targetIdx[k-lo] = i - lo
			i++
		} else {
			targetIdx[k-lo] = j - lo
			j++
		}
	}
	// swap all pairs
	for from, to := range findSwapPairs(targetIdx) {
		if from == to {
			continue
		}
		datas.Swap(from+lo, to+lo)
	}
}

func findSwapPairs(target []int) (swapPairs []int) {
	swapPairs = make([]int, len(target))
	for i := 0; i < len(target); i++ {
		swapPairs[i] = i
	}
	// here we can use path compression to improve performance
	for from, to := range target {
		if from == to {
			swapPairs[from] = to
			continue
		}
		for to < from {
			if to == swapPairs[to] {
				break
			}
			to = swapPairs[to]
		}
		swapPairs[from] = to
	}
	return swapPairs
}
