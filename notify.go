package main

import (
	"github.com/zikwall/go-hls/src/log"
	"os"
	"os/signal"
	"syscall"
)

func congratulations() {
	log.Info("Congratulations, the HLS server has been successfully launched")
}

func buildWaitNotifier() (func(), func()) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	wait := func() {
		// wait signal for the close application
		<-sig
	}

	stop := func() {
		// Send a signal to end the application
		sig <- syscall.SIGINT
	}

	return wait, stop
}
