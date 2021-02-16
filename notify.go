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

func buildWaitNotifier() chan os.Signal {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	return sig
}
