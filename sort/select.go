package sort

func SelectSort(data []int) {
	var m, l int
	l = len(data)

	for i := 0; i < l-1; i++ {
		m = i 
		for j := i+1; j < l; j++ {
			if data[j] < data[m] {
				m = j
			}
		}
		Swap(data, i, m)
	}
}