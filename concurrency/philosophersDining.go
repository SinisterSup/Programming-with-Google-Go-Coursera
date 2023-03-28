package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)

type ChopStick struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopStick
}

// philosopher eat method
func (p Philo) eat(philoNum int) {
	for i := 0; i < 3; i++ {
		//getting permission from host
		//blocking if host channel is full
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat: %d \n", philoNum)
		fmt.Println("Philosopher", philoNum, "is eating")
		time.Sleep(2000 * time.Millisecond)
		fmt.Printf("finshing eating: %d \n", philoNum)
		time.Sleep(750 * time.Millisecond)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		//relieve host channel
		//i.e., letting the host know eating is done
		<-host
	}
	wg.Done()
}

func main() {
	cSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philo{cSticks[i], cSticks[(i+1)%5]}
	}

	// Start dining
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(i + 1) // passing into eat is the Philosopher's ID
	}
	wg.Wait()

}
