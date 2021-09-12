package main

type Philosopher struct {
	nrEaten int
	eating  bool
	input   chan int
	output  chan int
	next    *Philosopher
	prev    *Philosopher
}

func eat(philosopher *Philosopher, fork1 *Fork, fork2 *Fork) {
	for {
		if !philosopher.prev.eating && !philosopher.next.eating {
			lock1.Lock()
			philosopher.eating = true
			fork1.input <- 1
			fork2.input <- 1
			philosopher.nrEaten++
			<-fork1.output
			<-fork2.output
			philosopher.eating = false
			lock1.Unlock()
		}
		if !philosopher.prev.eating && !philosopher.next.eating {
			lock2.Lock()
			philosopher.eating = true
			fork1.input <- 1
			fork2.input <- 1
			philosopher.nrEaten++
			<-fork1.output
			<-fork2.output
			philosopher.eating = false
			lock2.Unlock()
		}
	}
}
