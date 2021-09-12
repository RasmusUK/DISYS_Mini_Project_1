package main

import (
	"fmt"
	"sync"
	"time"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

func main() {

	phi1 := new(Philosopher)
	phi2 := new(Philosopher)
	phi3 := new(Philosopher)
	phi4 := new(Philosopher)
	phi5 := new(Philosopher)
	phi1.prev = phi5
	phi1.next = phi2
	phi2.prev = phi1
	phi2.next = phi3
	phi3.prev = phi2
	phi3.next = phi4
	phi4.prev = phi3
	phi4.next = phi5
	phi5.prev = phi4
	phi5.next = phi1

	fork1 := new(Fork)
	fork2 := new(Fork)
	fork3 := new(Fork)
	fork4 := new(Fork)
	fork5 := new(Fork)

	fork1.inUse = false
	fork2.inUse = false
	fork3.inUse = false
	fork4.inUse = false
	fork5.inUse = false

	makeChanPhi(phi1)
	makeChanPhi(phi2)
	makeChanPhi(phi3)
	makeChanPhi(phi4)
	makeChanPhi(phi5)

	makeChanFork(fork1)
	makeChanFork(fork2)
	makeChanFork(fork3)
	makeChanFork(fork4)
	makeChanFork(fork5)

	go eat(phi1, fork1, fork5)
	go eat(phi2, fork1, fork2)
	go eat(phi3, fork2, fork3)
	go eat(phi4, fork3, fork4)
	go eat(phi5, fork4, fork5)

	for {
		go display(phi1, 1)
		go display(phi2, 2)
		go display(phi3, 3)
		go display(phi4, 4)
		go display(phi5, 5)
		time.Sleep(time.Millisecond * 1000)
	}

}

func display(philosopher *Philosopher, number int) {
	fmt.Println("Philosopher ", number, " Eaten: ", philosopher.nrEaten)
}

func makeChanPhi(philosopher *Philosopher) {
	philosopher.output = make(chan int)
	philosopher.input = make(chan int)
}

func makeChanFork(fork *Fork) {
	fork.output = make(chan int)
	fork.input = make(chan int)
}
