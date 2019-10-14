package insertion

import (
	"algs4go/interfaces"
	"algs4go/utils"
	"testing"
)

var input utils.MyInts

func init() {
	input = utils.MyIntns(100, 100)
}

func TestInsertion(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := Insertion(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

// FastInsertion implements fast insertion sort algorithm by reducing
// accessments to datas(dirty implementation, DO NOT USE)
func FastInsertion(datas interfaces.Sortable) interfaces.Sortable {
	for i := 1; i < datas.Len(); i++ {
		now := datas.(utils.MyInts)[i]
		index := i
		for j := i; j > 0; j-- {
			if now < datas.(utils.MyInts)[j-1] {
				datas.(utils.MyInts)[j] = datas.(utils.MyInts)[j-1]
				index--
			}
		}
		datas.(utils.MyInts)[index] = now
	}
	return datas
}

func TestFastInsertion(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := FastInsertion(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

func BenchmarkInsertion(b *testing.B) {
	datas := utils.MyIntns(100000, 100000)
	for i := 0; i < b.N; i++ {
		Insertion(datas)
	}
}
