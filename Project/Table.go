package main

var cn = make(chan int, 2)

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
	phi1.nr = 1
	phi2.nr = 2
	phi3.nr = 3
	phi4.nr = 4
	phi5.nr = 5
	fork1.nr = 1
	fork2.nr = 2
	fork3.nr = 3
	fork4.nr = 4
	fork5.nr = 5
	cn <- 1
	cn <- 1

	makeChanFork(fork1)
	makeChanFork(fork2)
	makeChanFork(fork3)
	makeChanFork(fork4)
	makeChanFork(fork5)

	go eat(phi1, fork1, fork2)
	go eat(phi2, fork2, fork3)
	go eat(phi3, fork3, fork4)
	go eat(phi4, fork4, fork5)
	go eat(phi5, fork5, fork1)

	go run(fork1)
	go run(fork2)
	go run(fork3)
	go run(fork4)
	go run(fork5)

	for {
	}
}

func makeChanFork(fork *Fork) {
	fork.output = make(chan int)
	fork.input = make(chan int)
}
