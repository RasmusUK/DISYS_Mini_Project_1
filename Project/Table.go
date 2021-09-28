package main

var cn = make(chan int, 2)

func main() {

	forks := BuildForks()
	philosophers := BuildPhilosophers(forks)
	cn <- 1
	cn <- 1

	go PrintStatus(philosophers, forks)
	for {
	}
}

func makeChanFork(fork *Fork) {
	fork.output = make(chan int)
	fork.input = make(chan int)
	fork.status = make(chan string, 1)
	fork.status <- fmt.Sprintf("Fork %d is in use and has been used %d times", fork.nr, fork.nrUsed)
}

func makeChanPhi(phi *Philosopher) {
	phi.status = make(chan string, 1)
	phi.status <- fmt.Sprintf("Philosopher %d is thinking and has eaten %d times", phi.nr, phi.nrEaten)
}

func BuildPhilosophers(forks []Fork) []Philosopher {
	philosophers := make([]Philosopher, 0)
	for i := 0; i < 5; i++ {
		var phi = new(Philosopher)
		makeChanPhi(phi)
		phi.nr = i
		philosophers = append(philosophers, *phi)
		index := i + 1
		if i%4 == 0 && i != 0 {
			index = 0
		}
		go eat(phi, &forks[i], &forks[index])
	}
	return philosophers
}

func BuildForks() []Fork {
	var forks = make([]Fork, 0)
	for i := 0; i < 5; i++ {
		var fork = new(Fork)
		makeChanFork(fork)
		fork.nr = i
		forks = append(forks, *fork)
		go run(fork)
	}
	return forks
}

func PrintStatus(philosophers []Philosopher, forks []Fork) {
	for {
		for _, phi := range philosophers {
			status := <-phi.status
			fmt.Println(status)
			phi.status <- status
		}

		fmt.Println("")

		for _, fork := range forks {
			status := <-fork.status
			fmt.Println(status)
			fork.status <- status
		}

		fmt.Println("")
		fmt.Println(".......")
		fmt.Println("")
		time.Sleep(time.Millisecond * 1000)
	}
}
