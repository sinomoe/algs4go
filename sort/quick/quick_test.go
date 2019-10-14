package quick

import (
	"algs4go/utils"
	"sort"
	"testing"
)

var input utils.MyInts
var benchInput utils.MyInts

func init() {
	input = utils.MyIntns(100, 100)
	benchInput = utils.MyIntns(100000, 100000)
}

func TestQuick(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := Quick(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

func BenchmarkQuick(b *testing.B) {
	b.StopTimer()
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		b.StartTimer()
		Quick(datas)
		b.StopTimer()
	}
}
func BenchmarkInternalQuick(b *testing.B) {
	b.StopTimer()
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		b.StartTimer()
		sort.Sort(datas)
		b.StopTimer()
	}
}
