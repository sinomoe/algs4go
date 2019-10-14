package utils

import "math/rand"

// MyInts is a alias of []int, we implemented its Sortable interface
type MyInts []int

func (a MyInts) Len() int           { return len(a) }
func (a MyInts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MyInts) Less(i, j int) bool { return a[i] < a[j] }
func (a MyInts) Copy() MyInts {
	var cp MyInts
	copy(cp, a)
	return cp
}

// MyIntns returns a MyInts slice which has num elems, n denotes range [0, n)
func MyIntns(num, n int) MyInts {
	myInts := MyInts{}
	for i := 0; i < num; i++ {
		myInts = append(myInts, rand.Intn(n))
	}
	return myInts
}
