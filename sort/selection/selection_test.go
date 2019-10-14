package selection

import (
	"algs4go/utils"
	"testing"
)

func TestSelection(t *testing.T) {
	input := utils.MyIntns(100, 100)
	out := Selection(input)
	if !utils.IsSorted(out) {
		t.Fatal("selection sort falied")
	}
}
