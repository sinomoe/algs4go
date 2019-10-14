package merge

import (
	"algs4go/utils"
	"testing"
)

var input utils.MyInts
var benchInput utils.MyInts

func init() {
	input = utils.MyIntns(100, 100)
	benchInput = utils.MyIntns(100000, 100000)
}

func TestFindSwapParis(t *testing.T) {
	input := []int{0, 4, 1, 5, 6, 2, 3, 7}
	ans := []int{0, 4, 4, 5, 6, 6, 7}
	res := findSwapPairs(input)
	for k, v := range ans {
		if v != res[k] {
			t.Error("swap error")
		}
	}
	t.Log(res)
}

func TestInnerMerge(t *testing.T) {
	input := utils.MyInts([]int{1, 3, 7, 8, 2, 5, 6, 9})
	merge(input, 0, 3, 7)
	if !utils.IsSorted(input) {
		t.Error("merge error")
	}
}

func TestMerge(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := Merge(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

func TestMergeBU(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := MergeBU(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

func BenchmarkMerge(b *testing.B) {
	b.StopTimer()
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		b.StartTimer()
		Merge(datas)
		b.StopTimer()
	}
}

func BenchmarkMergeBU(b *testing.B) {
	b.StopTimer()
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		b.StartTimer()
		MergeBU(datas)
		b.StopTimer()
	}
}
