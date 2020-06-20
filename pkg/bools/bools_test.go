package bools

import "testing"

func TestAllTrue(t *testing.T) {
	bools := Bools{false}
	txt := "The result should be "
	if result := bools.AllTrue(); result {
		t.Error(txt + "false.")
	}

	max := 100 // a random big number
	bools = Bools(make([]bool, max, max))
	for i := range bools {
		bools[i] = true
	}

	bools[max-1] = false
	if result := bools.AllTrue(); result {
		t.Error(txt + "false.")
	}

	bools[max-1] = true
	if result := bools.AllTrue(); !result {
		t.Error(txt + "ture.")
	}

}
