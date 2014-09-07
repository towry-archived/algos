package sort

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	// k is used to mark the last swapped index
	var tmp, l, k, p int

	l = len(data)
	p = 0
	k = p + 1

	// set first elem as pivot
	for j := p + 1; j < l; j++ {
		// if less than pivot
		if data[j] < data[p] {
			tmp = data[j]
			data[j] = data[k]
			data[k] = tmp
			k++
		}
	}
	// now exchange the pivot with the rightmost elem in the left part.
	// so the pivot is in it's final position.
	tmp = data[p]
	data[p] = data[k-1]
	// pivot now is at `k-1`
	data[k-1] = tmp

	// partition is done

	// quick sort left
	QuickSort(data[:(k-1)])
	// quick sort right
	QuickSort(data[k:])
}