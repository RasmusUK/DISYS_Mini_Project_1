package main

import (
	"fmt"
	"sync"
)

type Fork struct {
	nrUsed int
	inUse  bool
	nr     int
	input  chan int
	output chan int
	status chan string
	lock   sync.Mutex
}

func run(fork *Fork) {
	for {
		<-fork.input
		fork.inUse = true
		fork.nrUsed++
		updateStatus(fork)
		<-fork.input
		fork.inUse = false
		updateStatus(fork)
		fork.output <- 1
	}

}

func updateStatus(fork *Fork) {
	if len(fork.status) != 0 {
		<-fork.status
	}
	var status string
	if fork.inUse {
		status = fmt.Sprintf("Fork ", fork.nr, " is in use and has been used ", fork.nrUsed, " times")
	} else {
		status = fmt.Sprintf("Fork ", fork.nr, " is not in use and has been used ", fork.nrUsed, " times")
	}
	fork.status <- status
}
