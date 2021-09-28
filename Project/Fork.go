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
		go updateStatusFork(fork)
		<-fork.input
		fork.inUse = false
		go updateStatusFork(fork)
		fork.output <- 1
	}
}

func updateStatusFork(fork *Fork) {
	if len(fork.status) != 0 {
		<-fork.status
	}
	var status string
	if fork.inUse {
		status = fmt.Sprintf("Fork %d is in use and has been used %d times", fork.nr, fork.nrUsed)
	} else {
		status = fmt.Sprintf("Fork %d is not in use and has been used %d times", fork.nr, fork.nrUsed)
	}
	fork.status <- status
}
