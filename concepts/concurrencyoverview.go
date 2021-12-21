package concepts

import (
	"fmt"
	"math/rand"
	"time"
)

// FAN IN
// ch chan <-  interface{} means that we only write to channel. Write only channel
// ch <- chan interface{} means that we can only read from channel. Read only channel
func fanInProducer(ch chan <- int, name string)  {
	for {
		// generate random number
		n := rand.Intn(100)
		// sleep for some time to simulate work
		time.Sleep(time.Duration(n) * 2 * time.Millisecond)
		// write the generated random number to the channel
		fmt.Printf("%s -> %d\n", name, n)
		ch <- n
	}
	
}

func fanInConsumer(ch <- chan int, name string) {
	for n := range ch {
		fmt.Printf("%s <- %d\n", name, n)
	}
}

func FanInExample() {
	var n int

	chanA := make(chan  int)
	chanB := make(chan  int)
	chanC := make(chan  int)

	go fanInProducer(chanA, "Channel A")
	go fanInProducer(chanB, "Channel B")
	go fanInConsumer(chanC, "Channel C")

	for {
		select {
		case n = <- chanA:
			chanC <- n
		case n = <- chanB:
			chanC <- n
		}
	}


}

// FAN OUT
func fanOutProducer(ch chan <- int, name string)  {
	for {
		// generate random number
		n := rand.Intn(100)
		// sleep for some time to simulate work
		time.Sleep(time.Duration(n) * 2 * time.Millisecond)
		// write the generated random number to the channel
		fmt.Printf("%s -> %d\n", name, n)
		ch <- n
	}

}

func fanOutConsumer(ch <- chan int, name string) {
	for n := range ch {
		fmt.Printf("%s <- %d\n", name, n)
	}
}

func FanOutExample() {
	var n int

	chanA := make(chan  int)
	chanB := make(chan  int)
	chanC := make(chan  int)

	go fanOutProducer(chanA, "Channel A")
	go fanOutConsumer(chanB, "Channel B")
	go fanOutConsumer(chanC, "Channel C")

	//for {
	//	select {
	//	case n = <- chanA:
	//		chanC <- n
	//	case n = <- chanB:
	//		chanC <- n
	//	}
	for n = range chanA {
		if n < 50 {
			chanB <- n
		} else {
			chanC <- n
		}
	}


}