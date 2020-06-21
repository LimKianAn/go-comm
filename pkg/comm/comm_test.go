package comm

import (
	"math"
	"testing"
	"time"

	"github.com/LimKianAn/go-comm/pkg/msg"
	"github.com/LimKianAn/go-comm/pkg/randsec"
)

func TestCycSendID(t *testing.T) {

	comm := New(1) // single tx
	ID := 0        // single tx
	sec := 1       // The duration is 1 second.
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
}

func TestCycSendDuration(t *testing.T) {
	tolerance := 0.05

	comm := New(1) // single tx
	ID := 0        // single tx
	sec := 1       // The duration is 1 second.
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

func TestReceive(t *testing.T) {
	txN := 100 // number of tx
	comm := New(txN)
	for i := 0; i < txN; i++ {
		go comm.CycSend(i, randsec.Get(15))
	}

	msgSlice := []*msg.Msg{} // placeholder to save all received messages
	comm.Receive(func(m *msg.Msg) {
		msgSlice = append(msgSlice, m)
	})

	lastID := msgSlice[len(msgSlice)-1].ID // tx ID of the last message
	counter := 0
	for _, e := range msgSlice {
		if e.ID == lastID {
			counter++
		}
	}

	if counter != 1 {
		t.Error("The last tx ID should have only occured once.")
	}

	for i := 0; i < txN; i++ {
		counter := 0
		for _, e := range msgSlice {
			if e.ID == i {
				counter++
			}
		}

		if counter == 0 {
			t.Errorf("Each tx ID should have occured at least once, but tx ID %v has never occured.", i)
		}
	}
}
