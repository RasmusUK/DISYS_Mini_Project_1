package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

var phi1 = new(Philosopher)
var phi2 = new(Philosopher)
var phi3 = new(Philosopher)
var phi4 = new(Philosopher)
var phi5 = new(Philosopher)

var fork1 = new(Fork)
var fork2 = new(Fork)
var fork3 = new(Fork)
var fork4 = new(Fork)
var fork5 = new(Fork)

func main() {

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

	go run(fork1)
	go run(fork2)
	go run(fork3)
	go run(fork4)
	go run(fork5)

	go display()

	for {
	}

}

func display() {
	for {
		fmt.Printf("\r%s", toString())
		time.Sleep(time.Millisecond * 1000)
	}
}

func toString() string {
	var line strings.Builder
	line.WriteString(fmt.Sprintf("\nPhilopsher %d Number of times eaten: %d", 1, phi1.nrEaten))
	line.WriteString(fmt.Sprintf("\nPhilopsher %d Number of times eaten: %d", 2, phi2.nrEaten))
	line.WriteString(fmt.Sprintf("\nPhilopsher %d Number of times eaten: %d", 3, phi3.nrEaten))
	line.WriteString(fmt.Sprintf("\nPhilopsher %d Number of times eaten: %d", 4, phi4.nrEaten))
	line.WriteString(fmt.Sprintf("\nPhilopsher %d Number of times eaten: %d", 5, phi5.nrEaten))
	line.WriteString(fmt.Sprintf("\nFork %d Number of times used: %d", 1, fork1.nrUsed))
	line.WriteString(fmt.Sprintf("\nFork %d Number of times used: %d", 2, fork2.nrUsed))
	line.WriteString(fmt.Sprintf("\nFork %d Number of times used: %d", 3, fork3.nrUsed))
	line.WriteString(fmt.Sprintf("\nFork %d Number of times used: %d", 4, fork4.nrUsed))
	line.WriteString(fmt.Sprintf("\nFork %d Number of times used: %d", 5, fork5.nrUsed))
	return line.String()
}

func makeChanPhi(philosopher *Philosopher) {
	philosopher.output = make(chan int)
	philosopher.input = make(chan int)
}

func makeChanFork(fork *Fork) {
	fork.output = make(chan int)
	fork.input = make(chan int)
}
