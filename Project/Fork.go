package main

import "sync"

type Fork struct {
	nrUsed int
	inUse  bool
	mutex  sync.Mutex
	input  chan int
	output chan int
}

func run(fork *Fork, phi1 *Philosopher, phi2 *Philosopher) {
	<-phi1.input
	fork.mutex.Lock()
	fork.inUse = true
	fork.nrUsed++
	phi1.input <- 1
	<-phi1.output
	fork.inUse = false
	fork.mutex.Unlock()

	<-phi2.input
	fork.mutex.Lock()
	fork.inUse = true
	fork.nrUsed++
	phi2.input <- 1
	<-phi2.output
	fork.inUse = false
	fork.mutex.Unlock()
}
