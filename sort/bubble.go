package sort

func BubbleSort(data []int) {
	swapped := true
	l := len(data)

	for swapped {
		swapped = false

		for i := 1; i < l; i++ {
			if data[i] < data[i-1] {
				Swap(data, i, i-1)
				swapped = true
			}
		}
	}
}