package main

import (
	"github.com/inf0rmatiker/go-playground/internal/examples/concurrency"
)

func main() {
	// fmt.Printf("Printf message\n")
	// log.Info("Logged message\n")

	concurrency.DoWorkPipeline()

}
