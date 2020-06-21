package comm

import (
	"time"

	"github.com/LimKianAn/go-comm/pkg/bools"
	"github.com/LimKianAn/go-comm/pkg/msg"
)

type Comm struct { // communication
	txN  int           // number of tx
	msg  chan *msg.Msg // channel of pointers
	done chan struct{} // only for signaling the end of communication
}

func New(txN int) *Comm { // allocates and initializes the channels
	return &Comm{txN, make(chan *msg.Msg), make(chan struct{})}
}

func (comm *Comm) CycSend(ID int, d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			comm.msg <- msg.New(ID)
		case <-comm.done: // Rx has received a complete set.
			return
		default: // non-blocking
		}
	}
}

func (comm *Comm) Receive(f msg.Process) { // txN means how many tx in toal.
	bools := bools.Bools(make([]bool, comm.txN, comm.txN)) // casts to type Bools
	for {
		m := <-comm.msg
		f(m)
		if bools[m.ID] = true; bools.AllTrue() { // checks if the rx has received a complete set
			comm.done <- struct{}{}
			return
		}
	}
}
