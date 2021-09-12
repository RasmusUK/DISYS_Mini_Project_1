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
			fork1.inUse = true
			fork2.inUse = true
			fork1.nrUsed++
			fork2.nrUsed++
			philosopher.nrEaten++
			fork1.inUse = false
			fork2.inUse = false
			philosopher.eating = false
			lock1.Unlock()
		}
		if !philosopher.prev.eating && !philosopher.next.eating {
			lock2.Lock()
			philosopher.eating = true
			fork1.inUse = true
			fork2.inUse = true
			fork1.nrUsed++
			fork2.nrUsed++
			philosopher.nrEaten++
			fork1.inUse = false
			fork2.inUse = false
			philosopher.eating = false
			lock2.Unlock()
		}
	}
}
