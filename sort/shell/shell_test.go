package shell

import (
	"algs4go/utils"
	"testing"
)

var input utils.MyInts

func init() {
	input = utils.MyIntns(100, 100)
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
	datas := utils.MyIntns(100000, 100000)
	for i := 0; i < b.N; i++ {
		Shell(datas)
	}
}
