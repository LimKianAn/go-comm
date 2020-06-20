package msg

import (
	"fmt"
	"time"
)

type Msg struct {
	ID                int
	TimeStampUnixNano int64
}

func Make(txID int) *Msg {
	return &Msg{txID, time.Now().UnixNano()}
}

func (m *Msg) String() string {
	return fmt.Sprintf("%2v (ID), %v", m.ID, time.Unix(0, m.TimeStampUnixNano))
}

type Process func(*Msg)
