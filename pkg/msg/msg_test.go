package msg

import (
	"fmt"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	n := time.Now().UnixNano()
	m := Msg{9, n}
	result := m.String()
	expectation := fmt.Sprintf("%2v (ID), %v", m.ID, time.Unix(0, n))
	if result != expectation {
		t.Errorf("The string should be %v, but we got %v", expectation, result)
	}
}
