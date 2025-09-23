package examples

import (
	"fmt"
	"time"
)

// Channels examples

// Adds an int to the channel. Blocks until the other side is ready to receive.
// This function should be executed as a goroutine, otherwise you'll end up in
// a deadlock.
func addValueToChannel(value int, c chan int) {
	fmt.Printf("Requested to add %d to channel\n", value)
	c <- value
	fmt.Printf("Added value %d to channel\n", value)
}

func BasicExample() {
	fmt.Println("-------- BasicExample --------")

	ch := make(chan int)
	// kick off goroutine, which will block on adding value until this func
	// is ready to receive.
	go addValueToChannel(42, ch)
	time.Sleep(1 * time.Second)
	go addValueToChannel(43, ch)
	time.Sleep(1 * time.Second)

	what := <-ch // receive value from channel
	fmt.Printf("what=%d\n", what)
	time.Sleep(1 * time.Second)
	what = <-ch
	fmt.Printf("what=%d\n", what)
	time.Sleep(1 * time.Second)

	fmt.Println("------------------------------")

	/* Output:
	-------- BasicExample --------
	Requested to add 42 to channel
	Requested to add 43 to channel
	what=42
	Added value 42 to channel
	what=43
	Added value 43 to channel
	------------------------------

	The time.Sleep()s are important, because otherwise this function can
	finish before the last goroutine has had a chance to finish.
	Also, both goroutines are executed in parallel, so this ensures ordering
	of the print statements.
	*/
}

func addToBufferedChannel(c chan int) {
	for i := 1; i <= 300; i++ {
		c <- i
		if i != 0 && i%100 == 0 {
			fmt.Printf("Added %d items to channel\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}

func BufferedChannels() {
	ch := make(chan int, 100)
	go addToBufferedChannel(ch)
	for i := range ch {
		fmt.Printf("i=%d\n", i)
	}
}
