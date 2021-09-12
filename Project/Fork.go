package main

type Fork struct {
	nrUsed int
	inUse  bool
	input  chan int
	output chan int
}

func run(fork *Fork, phi1 *Philosopher, phi2 *Philosopher) {
	for {
	}
}
