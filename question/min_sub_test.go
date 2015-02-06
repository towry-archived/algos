//
// (c) 2015 Towry Wang
//

package question

import "testing"

func TestMinSub(t *testing.T) {
	input_test := "aabcadbbbcca"
	expecting := "bcad"

	input_test2 := "k2www4oio42wk"
	expecting2 := "io42wk"

	output := MinSub(input_test)

	if output != expecting {
		t.Fail()
	}

	output := MinSub(input_test2)

	if output != expecting2 {
		t.Fail()
	}
}