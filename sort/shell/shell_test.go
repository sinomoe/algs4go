package shell

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

func TestShell(t *testing.T) {
	datas := make(utils.MyInts, len(input))
	copy(datas, input)

	out := Shell(datas)
	if !utils.IsSorted(out) {
		t.Fatal("insertion sort falied")
	}
}

func BenchmarkShell(b *testing.B) {
	b.StopTimer()
	datas := make(utils.MyInts, len(benchInput))
	for i := 0; i < b.N; i++ {
		copy(datas, benchInput)
		b.StartTimer()
		Shell(datas)
		b.StopTimer()
	}
}
