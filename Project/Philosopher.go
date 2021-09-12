package main

import (
	"sync"
)

type Philosopher struct {
	nrEaten int
	eating  bool
	input   chan int
	output  chan int
	next    *Philosopher
	prev    *Philosopher
}

func eat(philosopher *Philosopher) {
	philosopher.input <- 1
	philosopher.mutex.Lock()
	philosopher.eating = true
	philosopher.nrEaten++
	philosopher.output <- 1
	philosopher.eating = false
	philosopher.mutex.Unlock()
}
