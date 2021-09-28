package main

import (
	"time"
)

type Philosopher struct {
	nrEaten int
	nr      int
	eating  bool
}

func eat(philosopher *Philosopher, fork1 *Fork, fork2 *Fork) {
	for {
		<-cn
		fork1.lock.Lock()
		fork2.lock.Lock()
		fork1.input <- 1
		fork2.input <- 1
		philosopher.eating = true
		philosopher.nrEaten++
		time.Sleep(time.Millisecond * 500)
		fork1.input <- 1
		fork2.input <- 1
		<-fork1.output
		<-fork2.output
		philosopher.eating = false
		fork1.lock.Unlock()
		fork2.lock.Unlock()
		cn <- 1
	}
}
