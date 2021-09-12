package main

type Fork struct {
	nrUsed int
	inUse  bool
	input  chan int
	output chan int
}

func run(fork *Fork) {
	for {
		<-fork.input
		fork.inUse = true
		fork.nrUsed++
		fork.inUse = false
		fork.output <- 1
	}

}
