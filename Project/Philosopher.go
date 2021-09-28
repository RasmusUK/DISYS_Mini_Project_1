package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	nrEaten int
	nr      int
	eating  bool
	status  chan string
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
		go updateStatusPhi(philosopher)
		time.Sleep(time.Millisecond * 500)
		fork1.input <- 1
		fork2.input <- 1
		<-fork1.output
		<-fork2.output
		philosopher.eating = false
		go updateStatusPhi(philosopher)
		fork1.lock.Unlock()
		fork2.lock.Unlock()
		cn <- 1
	}
}

func updateStatusPhi(phi *Philosopher) {
	if len(phi.status) != 0 {
		<-phi.status
	}
	var status string
	if phi.eating {
		status = fmt.Sprintf("Philosopher %d is eating and has eaten %d times", phi.nr, phi.nrEaten)
	} else {
		status = fmt.Sprintf("Philosopher %d is thinking and has eaten %d times", phi.nr, phi.nrEaten)
	}
	phi.status <- status
}
