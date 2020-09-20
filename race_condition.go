/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/

/*
Race condition: This is a "data race" type race condition. We have a single variable x which is being written/updated  by
two separate a go routines and read by one go routine. Since the scheduling of the go routines are non deterministic the
final value of x is also non-deterministic and hence resulting in a data race condition.
*/
package main

import (
	"fmt"
	"time"
)

var x int

func main() {
	x = 0
	fmt.Println("The initial value of X is ", x)
	go increase()
	go decrease()
	go getX()
	// the above go routines won't execute if the main go routine finishes before hand so we are adding a delay/sleep
	time.Sleep(time.Second)
}

func increase() {
	fmt.Println("Increasing X by 1 ", x+1)
}

func decrease() {
	fmt.Println("Decreasing X by 1 ", x-1)
}

func getX() {
	fmt.Println("The final value of X is ", x)
}
