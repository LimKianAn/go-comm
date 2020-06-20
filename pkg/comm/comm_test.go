package comm

import (
	"math"
	"testing"
	"time"

	"github.com/LimKianAn/go-comm/msg"
)

func TestCycSendID(t *testing.T) {

	comm := Make(1) // single tx
	ID := 0         // single tx
	sec := 1        // The duration is 1 second.
	dur := time.Duration(sec) * time.Second

	receivedMsg := &msg.Msg{}
	// start := time.Now() // start of the communication
	go comm.CycSend(ID, dur)
	comm.Receive(func(m *msg.Msg) {
		receivedMsg = m //
	})

	receivedID := receivedMsg.ID
	if receivedID != ID {
		t.Errorf("The ID should be %v, but we got %v", ID, receivedID)
	}

	// if diff := time.Unix(0, receivedMsg.TimeStampUnixNano).Sub(start); !withinTolerance(diff, dur, tolerance) {
	// 	t.Errorf("The duration should be %v, but we got %v.", dur, diff)
	// }
}

func TestCycSendDuration(t *testing.T) {
	tolerance := 0.05

	comm := Make(1) // single tx
	ID := 0         // single tx
	sec := 1        // The duration is 1 second.
	dur := time.Duration(sec) * time.Second

	go comm.CycSend(ID, dur)

	n := 10 // number of durations
	for i := 0; i < n; i++ {
		msgA, msgB := <-comm.msg, <-comm.msg
		diff := time.Duration(msgB.TimeStampUnixNano - msgA.TimeStampUnixNano)
		if !withinTolerance(diff, dur, tolerance) {
			t.Errorf("The duration should be %v, but we got %v.", dur, diff)
		}
	}
}

func withinTolerance(diff, base time.Duration, tolerance float64) bool {
	return math.Abs(float64(diff-base)) < tolerance*float64(base)
}

// func isBetween(start, timepoint, end time.Time) bool {
// 	return timepoint.After(start) && timepoint.Before(start)
// }
