//
// (c) 2015 Towry Wang
//

package question

//# http://ask.julyedu.com/question/141

func MinSub(a string) string {
	var last_new, first_new, count int
	var d string

	// find last_new, and total count
	for i, _ := range a {
		if indexOf(a, a[i], 0, i) == -1 {
			last_new = i
			count += 1
		}
	}

	// in reverse order, start from last_new
	tmp := last_new - 1
	count = count - 1

	for tmp >= 0 {
		if a[tmp] == a[last_new] {
			first_new = last_new
			break
		}

		if indexOf(a, a[tmp], tmp+1, last_new) == -1 {
			count -= 1
		}

		if count == 0 {
			first_new = tmp
			break
		}

		tmp -= 1
	}
	
	d = a[first_new:last_new+1]
	
	return d
}

// search char in string between m - n
// base index is 0
// n is exclusive
func indexOf(s string, k uint8, m, n int) int {
	var i int

	i = m 
	for i < n {
		if (s[i] == k) {
			return i
		}
		i++
	}

	return -1
}
