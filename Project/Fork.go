package main

import "sync"

type Fork struct {
	nrUsed int
	inUse  bool
	nr     int
	input  chan int
	output chan int
	lock   sync.Mutex
}

func run(fork *Fork) {
	for {
		<-fork.input
		fork.inUse = true
		fork.nrUsed++
		<-fork.input
		fork.inUse = false
		fork.output <- 1
	}

}
