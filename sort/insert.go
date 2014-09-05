package sort

func InsertSort(data []int){
	l := len(data)

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if data[j] > data[i] {
				Swap(data, i, j)
			}
		}
	}
}