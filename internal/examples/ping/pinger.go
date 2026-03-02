package ping

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	probing "github.com/prometheus-community/pro-bing"
)

// Pinger interface to allow for testing of CheckConnectivity without sending real pings.
type PingerIface interface {
	Ping(context.Context, string, string, *log.Logger) error
}

type DefaultPinger struct{}

// First version of the Ping method that use an external 'success' channel to signal when to stop
// the pinger after receiving the first successful ping reply.
func (d *DefaultPinger) Ping(ctx context.Context, iface, addr string, logger *log.Logger) error {
	// Channel to signal first successful ping.
	success := make(chan bool)

	pinger, err := probing.NewPinger(addr)
	if err != nil {
		return err
	}
	pinger.SetPrivileged(true)
	pinger.SetLogger(logger)
	pinger.Interval = 1 * time.Second
	pinger.Count = 0
	pinger.InterfaceName = iface

	pinger.OnRecv = func(pkt *probing.Packet) {
		defer close(success) // sender closes the channel
		logger.Infof("Received ping reply from %s on interface %s: time=%v", pkt.IPAddr, iface, pkt.Rtt)
		success <- true
	}

	pinger.OnSend = func(pkt *probing.Packet) {
		logger.Infof("Sending ping...")
	}

	defer pinger.Stop()
	go pinger.Run()
	select {
	case <-success:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Second version of the Ping method that uses a parent context to control the pinger,
// through the pinger.RunWithContext() method, instead of a separate channel to signal when to stop the pinger.
func (d *DefaultPinger) Ping2(ctx context.Context, iface, addr string, logger *log.Logger) error {
	pinger, err := probing.NewPinger(addr)
	if err != nil {
		return err
	}
	pinger.SetPrivileged(true)
	pinger.SetLogger(logger)
	pinger.Interval = 1 * time.Second
	pinger.Count = 0
	pinger.InterfaceName = iface

	// Instead of manually controlling the pinger with a channel, use the parent context to stop the pinger,
	// and invoke pinger.Stop() if we successfully receive a ping reply.
	pinger.OnRecv = func(pkt *probing.Packet) {
		logger.Infof("Received ping reply from %s on interface %s: time=%v", pkt.IPAddr, iface, pkt.Rtt)
		pinger.Stop() // stop after first reply
	}
	pinger.OnSend = func(pkt *probing.Packet) {
		logger.Infof("Sending ping...")
	}

	return pinger.RunWithContext(ctx)
}
