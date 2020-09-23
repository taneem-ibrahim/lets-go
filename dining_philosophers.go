/*
Implement the dining philosopher’s problem with the following constraints/modifications.

- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
  on a line by itself, where <number> is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
  on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

// chopstick is safe to use concurrently (only one go routine can access it at a time)
type chopstick struct {
	mux sync.Mutex
}

type philosopher struct {
	leftChopStick, rightChopStick *chopstick
	idNumber                      int
	eatCount                      int
}

// note: function receiver type is philosopher
func (p philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for p.eatCount < 3 {
		p.leftChopStick.mux.Lock()
		p.rightChopStick.mux.Lock()
		fmt.Println("Philosopher #", p.idNumber, " is starting to eat.")
		fmt.Println("Philosopher #", p.idNumber, " is finishing eating.")
		p.rightChopStick.mux.Unlock()
		p.leftChopStick.mux.Unlock()
		p.eatCount++
	}
}

func main() {
	//create 5 chopsticks slice
	chops := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(chopstick)
	}
	//create 5 philosophers: notice the right chopstick is modulo 5 so that the 5th philosopher gets chopstick # 0
	//also assign an id number to each philosopher to keep track of them eating
	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{chops[i], chops[(i+1)%5], i, 0}
	}
	// eat
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg)
	}
	wg.Wait()
}
