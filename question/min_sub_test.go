//
// (c) 2015 Towry Wang
//

package question

import "testing"

func TestMinSub(t *testing.T) {
	var output string

	input_test := "aabcadbbbcca"
	expecting := "bcad"

	input_test2 := "k2www4oio42wk"
	expecting2 := "io42wk"

	input_test3 := "abcdefghijk"
	expecting3 := "abcdefghijk"

	output = MinSub(input_test)

	if output != expecting {
		t.Log("fail1:%s", output)
		t.Fail()
	}

	output = MinSub(input_test2)

	if output != expecting2 {
		t.Log("fail2")
		t.Fail()
	}

	output = MinSub(input_test3)

	if output != expecting3 {
		t.Log("fail3")
		t.Fail()
	}
}
