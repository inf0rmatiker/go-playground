package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inf0rmatiker/go-playground/internal/examples/ping"

	log "github.com/sirupsen/logrus"
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

	mainCtx, cancel := context.WithCancel(context.Background())
	registerSigHandler(cancel)
	defer cancel()

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <interface> <address>", os.Args[0])
	}
	iface, addr := os.Args[1], os.Args[2]

	pingerTimeout := 5 * time.Second // time out after 5 seconds
	pingerCtx, pCancel := context.WithTimeout(mainCtx, pingerTimeout)
	defer pCancel()
	pinger := ping.DefaultPinger{}
	err := pinger.Ping2(pingerCtx, iface, addr, log.StandardLogger())
	if err != nil {
		log.Error(err)
	}
}
