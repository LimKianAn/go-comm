package randsec

import (
	"testing"
)

func TestGet(t *testing.T) {
	max := 10
	for i := 0; i < max*max*max; i++ {
		if v := int(Get(max).Seconds()); v > 10 {
			t.Errorf("The duration should be equal to or less than 10 seconds (<=), we but got %v", v)
		}
	}
}
