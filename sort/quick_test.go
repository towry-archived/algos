package sort

import "testing"

func Test_QuickSort(t *testing.T) {
	data := []int {5, 3, 12, 9, 7, 23, 18}

	QuickSort(data)

	if data[0] != 3 {
		t.Fail()
	}

	Print(data, "quick")
}