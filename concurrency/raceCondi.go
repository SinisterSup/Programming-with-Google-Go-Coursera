// main() executes two anonymous functions. Keyword 'go' makes them run
// concurrently, in goroutines.
// Go runtime scheduler schedules goroutines in a non-deterministic way so
// there are multiple possible interleavings - each time we run the application
// scheduler might change the order of their execution.
// Sometimes function which increments x will be executed first while sometimes the function which prints out its value.
// Because both functions in goroutines are communicating through variable x
// there is a chance of race condition - the order of accessing variable x is non-deterministic.
// If we run the application multiple times, it will sometimes print 0 and sometimes 1.
//
// To test this, you can run this application in loop.
// Please use the following command in terminal (bash):
// for i in (1, 1, 100) do go run racecondition.go
//
// This randomness in output proves there is a race condition.

package concurrency

import (
	"fmt"
	"time"
)

var x int

// One goroutine implements an increment of the variable.
// The other one prints the value of the variable on screen.
func main() {
	go func() {
		x++
	}()

	go func() {
		fmt.Println("Value of x: ", x)
	}()

	// prevent program termination before goroutines return
	time.Sleep(1 * time.Second)
}

//Race Condition is when multiple threads in a concurrent running program accessing the same shared variable and attempt to modify the variable at the same time.
