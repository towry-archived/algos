package sort

import "testing"

func Test_InsertSort(t *testing.T) {
	data := []int {5, 3, 12, 9, 7, 23, 18}

	InsertSort(data)

	if data[0] != 3 {
		t.Fail()
	}

	Print(data, "insert")
}