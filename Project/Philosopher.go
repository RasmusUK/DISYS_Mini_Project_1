package main

import (
	"fmt"
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
		fmt.Println("Phi", philosopher.nr, "allowed")
		fork1.lock.Lock()
		fork1.inUse = true
		fmt.Println("Phi", philosopher.nr, "picked up fork", fork1.nr)
		fork2.lock.Lock()
		fork2.inUse = true
		fmt.Println("Phi", philosopher.nr, "picked up fork", fork2.nr)
		fmt.Println("Phi", philosopher.nr, "eating")
		fork1.input <- 1
		fork2.input <- 1
		philosopher.eating = true
		philosopher.nrEaten++
		time.Sleep(time.Millisecond * 1000)
		fork1.input <- 1
		fork2.input <- 1
		<-fork1.output
		<-fork2.output
		philosopher.eating = false
		fmt.Println("Phi", philosopher.nr, "done eating. Eaten:", philosopher.nrEaten, "times")
		fork1.lock.Unlock()
		fork1.inUse = false
		fmt.Println("Phi", philosopher.nr, "put down fork", fork1.nr)
		fork2.lock.Unlock()
		fork2.inUse = false
		fmt.Println("Phi", philosopher.nr, "put down fork", fork2.nr)
		cn <- 1
	}
}
