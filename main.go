package main

import (
	"context"
	"fmt"

	"github.com/go-ping/ping"
)

func Run(ctx context.Context, target string, command string, expectedOutput string, username string, password string) (bool, string) {
	// create ping
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return false, err.Error()
	}

	pinger.Count = 1
	doneChan := make(chan error, 1)

	// run ping
	go func() {
		defer close(doneChan)
		doneChan <- pinger.Run()
	}()

	// handle ping output
	select {
	case err := <-doneChan:
		if err != nil {
			// error encoutnered when attempting to ping
			return false, fmt.Sprintf("Encounter error: %s", err)
		}

		stats := pinger.Statistics()

		if stats.PacketLoss == 0 {
			// successful ping
			return true, ""
		} else {
			// packet loss not 0
			return false, fmt.Sprintf("Packet loss: %.2f%%", stats.PacketLoss)
		}

	case <-ctx.Done():
		pinger.Stop()

		// context exceed; timeout
		return false, fmt.Sprintf("Timeout exceeded; err: %s", ctx.Err())
	}
}
