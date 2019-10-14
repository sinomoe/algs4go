package selection

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

func TestSelection(t *testing.T) {
	input := utils.MyIntns(100, 100)
	out := Selection(input)
	if !utils.IsSorted(out) {
		t.Fatal("selection sort falied")
	}
}

func BenchmarkSelection(b *testing.B) {
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		Selection(datas)
	}
}
