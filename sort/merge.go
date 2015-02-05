package sort

// Let the segment be 1, 2, 4, 8 
func MergeSort(a []int, size int) {
	tmp := make([]int, size)
	seg := 1

	// division
	for seg < size {
		mergePass(a, tmp, seg, size)
		seg += seg
		mergePass(tmp, a, seg, size)
		seg += seg
	}
}

// When the segment is 1, every piece is sorted, merge two pieces is easy
// then when the segment is 2, merge two pieces together
func mergePass(a []int, b []int, seg, size int) {
	seg_start_ind := 0
	// suppose current segment is 1
	// 2, 3, 1, 4, 12
	for seg_start_ind <= size - 2 * seg {
		// suppose it's the first pass, here we only merge 2, 3
		merge(a, b, seg_start_ind, seg_start_ind + seg - 1, seg_start_ind + seg * 2 - 1)
		seg_start_ind += 2 * seg;
	}

	// merge the rest
	if (seg_start_ind + seg < size) {
		merge(a, b, seg_start_ind, seg_start_ind + seg - 1, size - 1)
	} else {
		// merge the rest
		for j := seg_start_ind; j < size; j++ {
			b[j] = a[j]
		}
	}
}

// merge 
func merge(a, b []int, low, mid, high int) {
	var k int
	var begin1 int
	var begin2 int
	var end1 int
	var end2 int

	k = low
	begin1 = low
	begin2 = mid + 1
	end1 = mid 
	end2 = high

	for k <= high {
		if begin1 > end1 {
			b[k], k, begin2 = a[begin2], k+1, begin2+1
		} else if begin2 > end2 {
			b[k], k, begin1 = a[begin1], k+1, begin1+1
		} else {
			if a[begin1] <= a[begin2] {
				b[k], k, begin1 = a[begin1], k+1, begin1+1
			} else {
				b[k], k, begin2 = a[begin2], k+1, begin2+1
			}
		}
	}
}