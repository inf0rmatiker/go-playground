package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inf0rmatiker/go-playground/internal/examples"
)

func registerSigHandler(ctxCancel context.CancelFunc) {

	// 1. Create a channel to receive OS signals. It should be buffered.
	sigChannel := make(chan os.Signal, 1)

	// 2. Register the channel to receive SIGINT and SIGTERM.
	// Use os.Interrupt for cross-platform compatibility instead of syscall.SIGINT.
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	// 3. Start a goroutine to wait for the signal.
	go func() {
		sig := <-sigChannel
		fmt.Printf("\nReceived signal: %s. Performing cleanup...\n", sig)

		// Cancel context
		ctxCancel()
	}()

}

func main() {

	_, cancel := context.WithCancel(context.Background())
	registerSigHandler(cancel)
	defer cancel()

	err := examples.ReadLines()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

}
