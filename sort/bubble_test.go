package sort

import "testing"

func Test_BubbleSort(t *testing.T) {
	data := []int {5, 3, 12, 9, 7, 23, 18}

	BubbleSort(data)

	if data[0] != 3 {
		t.Fail()
	}

	Print(data, "bubble")
}