package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LimKianAn/go-comm/pkg/comm"
	"github.com/LimKianAn/go-comm/pkg/msg"
	"github.com/LimKianAn/go-comm/pkg/randsec"
)

func main() {
	n := getTxN()
	comm := comm.Make(n)
	for i := 0; i < n; i++ {
		go comm.CycSend(i, randsec.Get(10))
	}

	comm.Receive(print)
}

func getTxN() int { // number of tx
	str := "Please enter a valid integer as the only argument."
	if len(os.Args) != 2 { // program name and the argument "number of tx"
		log.Fatalln(str)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		log.Fatalln(str)
	}

	return n
}

func print(m *msg.Msg) {
	fmt.Println(m)
}
