package sort

import "testing"

func TestSwap(t *testing.T) {
	data := []int {1, 2}

	Swap(data, 0, 1)

	if data[0] == 1 {
		t.Fail()
	}
}