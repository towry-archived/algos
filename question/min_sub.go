//
// (c) 2015 Towry Wang
//

package question

//# http://ask.julyedu.com/question/141

func MinSub(a string) string {
	var sub []string

	for i, _ := range a {
		if !contains(sub, string(a[i])) {
			sub = append(sub, string(a[i]))
		}
	}

	sub_len := len(sub)
	a_len := len(a)

	var min int
	var max int
	var pos int
	var bus []string

	for i := 0; i <= a_len - sub_len; i += 1 {
		for m := i; m < a_len; m += 1 {
			if !contains(bus, string(a[m])) {
				bus = append(bus, string(a[m]))
			}

			if len(bus) == sub_len {
				if min == 0 {
					min = max
					pos = i
				} else if min > max {
					min = max
					pos = i
				}
				break
			}

			max += 1
		}

		max = 0
		bus = nil
	}

	return a[pos:pos+min+1]
}

func contains(a []string, b string) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}

	return false
}
