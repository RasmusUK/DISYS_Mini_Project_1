package main

var cn = make(chan int, 2)

func main() {

	forks := BuildForks()
	philosophers := BuildPhilosophers(forks)
	cn <- 1
	cn <- 1

	/*
		go eat(phi1, fork1, fork2)
		go eat(phi2, fork2, fork3)
		go eat(phi3, fork3, fork4)
		go eat(phi4, fork4, fork5)
		go eat(phi5, fork5, fork1)
	*/
	go PrintStatus(philosophers, forks)
	for {
	}
}

func makeChanFork(fork *Fork) {
	fork.output = make(chan int)
	fork.input = make(chan int)
	fork.status = make(chan string, 1)
}
func BuildPhilosophers(forks []Fork) []Philosopher {
	philosophers := make([]Philosopher, 0)
	for i := 0; i < 5; i++ {
		var phi = new(Philosopher)
		phi.nr = i
		philosophers = append(philosophers, *phi)
		go eat(phi, &forks[i], &forks[((i)%4)+1])
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

/*
func PrintStatus(philosophers []Philosopher, forks []Fork){
	for{
		for _, phi := range philosophers{
			if phi.eating {
				fmt.Println("Philosopher ",phi.nr," is eating and has eaten ", phi.nrEaten, " times")
			}else{
				fmt.Println("Philosopher ",phi.nr," is thinking and has eaten ", phi.nrEaten, " times")
			}
		}

		fmt.Println("")

		for _, fork := range forks{
			if fork.inUse {
				fmt.Println("Fork ",fork.nr," is in use and has been used ", fork.nrUsed, " times")
			}else{
				fmt.Println("Fork ",fork.nr," is not in use and has been used ", fork.nrUsed, " times")
			}
		}

		fmt.Println("")
		fmt.Println(".......")
		fmt.Println("")
		time.Sleep(time.Millisecond * 1000)
	}
}

*/

func PrintStatus(philosophers []Philosopher, forks []Fork) {
	for {
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
