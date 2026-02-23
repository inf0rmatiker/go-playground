package concurrency

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func printWorker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range ch {
		fmt.Printf("printWorker: %d\n", x)
	}
}

// Invokes workers as goroutines, without needing to look at their results.
func DoWork() {
	wg := sync.WaitGroup{}
	evenChan, oddChan := make(chan int), make(chan int)
	wg.Add(2)
	go printWorker(evenChan, &wg)
	go printWorker(oddChan, &wg)

	for range 10 {
		newRand := rand.Int() % 10
		// fmt.Printf("DoWork: %d\n", newRand)
		if newRand%2 == 0 {
			evenChan <- newRand
		} else {
			oddChan <- newRand
		}
	}
	close(evenChan)
	close(oddChan)

	wg.Wait()
}

type result struct {
	original, sqrt float64
}

// Takes the square root of a float, and emits it on the channel.
func sqrt(ch chan result, wg *sync.WaitGroup, input float64) {
	defer wg.Done()
	ch <- result{original: input, sqrt: math.Sqrt(input)}
}

// Invokes workers as goroutines, gathering results at the end.
func DoWorkGatherResults() {
	originalNumbers := []float64{
		0.414, 0, 1, 4, 9, 12.345,
	}

	wg := sync.WaitGroup{}
	wg.Add(len(originalNumbers))
	results := make(chan result) // unbuffered channel

	// Kick off all the worker goroutines to calculate sqrts
	for _, number := range originalNumbers {
		go sqrt(results, &wg, number)
	}

	// Need a separate goroutine to wait for workers to finish, then close channel
	// so main goroutine finishes for loop.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results until channel is closed.
	for result := range results {
		fmt.Printf("original: %.4f, sqrt: %.4f\n", result.original, result.sqrt)
	}
}

func diff(a, b []float64) <-chan float64 {
	out := make(chan float64, 100)
	go func() {
		for i := range a {
			out <- a[i] - b[i]
		}
		close(out)
	}()
	return out
}

func square(in <-chan float64) <-chan float64 {
	out := make(chan float64, 100)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func sum(in <-chan float64) <-chan float64 {
	out := make(chan float64, 100)
	go func() {
		sum := 0.0
		for i := range in {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}

func rangedRandomFloats(min, max float64, count int) []float64 {
	floats := make([]float64, count)
	for i := range count {
		floats[i] = min + rand.NormFloat64()*(max-min)
	}
	return floats
}

// Invokes a series of goroutines in a pipeline, using channels to communicate from one stage to the next.
// We'll be calculating the Mean Squared Error (MSE) between two randomly-generated arrays of the same size.
func DoWorkPipeline() {
	size := 100_000
	actual := rangedRandomFloats(0, 100, size)
	predicted := rangedRandomFloats(0, 100, size)
	var mse float64

	begin := time.Now()
	stage1Chan := diff(actual, predicted)
	stage2Chan := square(stage1Chan)
	stage3Chan := sum(stage2Chan)

	sumResult := <-stage3Chan
	mse = sumResult / float64(len(actual))
	fmt.Printf("Pipeline MSE: %0.2f, time = %s\n", mse, time.Since(begin).String())

	actual = rangedRandomFloats(0, 100, size)
	predicted = rangedRandomFloats(0, 100, size)

	// Compare against just a normal, no goroutine implementation
	begin = time.Now()
	mse = 0.0
	for i, a := range actual {
		d := a - predicted[i]
		mse += d * d
	}
	mse /= float64(len(actual))
	fmt.Printf("Sequential MSE: %0.2f, time = %s\n", mse, time.Since(begin).String())
}
