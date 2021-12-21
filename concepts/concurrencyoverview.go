package concepts

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan <- int, name string)  {
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

func consumer(ch <- chan int, name string) {
	for n := range ch {
		fmt.Printf("%s <- %d\n", name, n)
	}
}

// FAN IN
// ch chan <-  interface{} means that we only write to channel. Write only channel
// ch <- chan interface{} means that we can only read from channel. Read only channel

func FanInExample() {
	var n int

	chanA := make(chan  int)
	chanB := make(chan  int)
	chanC := make(chan  int)

	go producer(chanA, "Channel A")
	go producer(chanB, "Channel B")
	go consumer(chanC, "Channel C")

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
func FanOutExample() {
	var n int

	chanA := make(chan  int)
	chanB := make(chan  int)
	chanC := make(chan  int)

	go producer(chanA, "Channel A")
	go consumer(chanB, "Channel B")
	go consumer(chanC, "Channel C")

	for n = range chanA {
		if n < 50 {
			chanB <- n
		} else {
			chanC <- n
		}
	}


}