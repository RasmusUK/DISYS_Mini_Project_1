package main

type Philosopher struct {
	nrEaten  int
	eating   bool
	adjForks []Fork
	input    chan int
	output   chan string
}
